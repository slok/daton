package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"

	"github.com/slok/daton/configuration"
	"github.com/slok/daton/utils"
)

const (
	// Deploy status constants
	StatusPending = "pending"
	StatusSuccess = "success"
	StatusError   = "error"
	StatusFailure = "failure"
)

// Deploy represents a deployment
type Deployment struct {
	Url           string      `json:"url"`
	Id            int64       `json:"id"`
	Sha           string      `json:"sha"`
	Ref           string      `json:"ref"` // Required always
	Task          string      `json:"task"`
	Payload       interface{} `json:"payload"`
	Environment   string      `json:"environment"`
	Description   string      `json:"description"`
	Creator       User        `json:"creator"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	StatusesUrl   string      `json:"statuses_url"`
	RepositoryUrl string      `json:"repository_url"`
	StatusId      int         `json:"status_id"`
	Namespace     string      `json:"-"`
}

func (d *Deployment) Save() error {
	// Set defaults
	if len(d.Task) <= 0 {
		d.Task = configuration.DefaultTask
	}
	if len(d.Environment) <= 0 {
		d.Environment = configuration.DefaultEnvironment
	}
	if d.Payload == nil {
		d.Payload = ""
	}

	// Generate the keys to insert the object id
	objectQueryKeys := []string{
		fmt.Sprintf(DeployQueryDbKeyFmt, d.Namespace, d.Environment),
		fmt.Sprintf(DeployQueryDbKeyFmt, d.Namespace, d.Task),
		fmt.Sprintf(DeployQueryDbKeyFmt, d.Namespace, d.Ref),
		fmt.Sprintf(DeployQueryDbKeyFmt, d.Namespace, d.Sha),
	}
	// Generate counter storer id
	deployCounterKey := fmt.Sprintf(DeployCounterKeyFmt, d.Namespace)

	// Start db stuff
	db, _ := GetBoltDb()
	err := db.Conn.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(DeployBucketDbKey))
		if err != nil {
			return err
		}

		// Get the last id and save + 1
		id := bucket.Get([]byte(deployCounterKey))
		var intId int64
		// If doens't exists means that is the first time for the counter
		if id == nil {
			intId = 0
		} else {
			intId, err = strconv.ParseInt(string(id), 10, 0)
		}
		intId = intId + 1
		err = bucket.Put([]byte(deployCounterKey), []byte(strconv.FormatInt(intId, 10)))
		if err != nil {
			return err
		}
		objectKey := fmt.Sprintf(DeployObjectDbKeyFmt, d.Namespace, intId)
		d.Id = intId

		// Insert deploy
		// Serialize to json
		b, err := json.Marshal(d)
		if err != nil {
			log.Errorf("Error serializing deployment: %v", err)
			return err
		}

		err = bucket.Put([]byte(objectKey), b)
		if err != nil {
			return err
		}

		// Add deploy to query keys
		for _, v := range objectQueryKeys {
			// Get the json (int array)
			deployListJson := bucket.Get([]byte(v))
			deployList := []int64{}
			if deployListJson != nil {
				err = json.Unmarshal(deployListJson, &deployList)
				if err != nil {
					return err
				}
			}

			// Append the new id
			deployList = append(deployList, intId)
			deployListJson, err = json.Marshal(deployList)
			if err != nil {
				return err
			}
			err = bucket.Put([]byte(v), deployListJson)
			if err != nil {
				return err
			}
		}

		log.WithFields(log.Fields{
			"namespace": d.Namespace,
			"id":        d.Id,
		}).Debug("Deployment inserted on database")
		return nil
	})

	if err != nil {
		log.Errorf("Error inserting deployment on database: %v", err)
		return err
	}

	return nil
}

func ListDeploymentsAsJson(namespace string) ([][]byte, error) {
	d := [][]byte{}
	//_ = "breakpoint"
	db, _ := GetBoltDb()
	err := db.Conn.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(DeployBucketDbKey))
		if bucket == nil {
			// No deployments
			return nil
		}

		c := bucket.Cursor()

		prefix := []byte(fmt.Sprintf(DeployObjectListDbKeyFmt, namespace))

		for k, v := c.Seek(prefix); bytes.HasPrefix(k, prefix); k, v = c.Next() {
			d = append(d, v)
		}
		return nil
	})

	log.WithFields(log.Fields{
		"namespace": namespace,
		"length":    len(d),
		"type":      "json",
	}).Debug("Deployments retrieved from database")
	return d, err
}

func ListDeployments(namespace string) ([]Deployment, error) {
	dJson, err := ListDeploymentsAsJson(namespace)
	if err != nil {
		return nil, err
	}
	ds := []Deployment{}
	for _, i := range dJson {
		d := &Deployment{}
		err := json.Unmarshal(i, d)
		if err != nil {
			return nil, err
		}

		// Create the dynamic urls
		d.Url = utils.ApiUrl(fmt.Sprintf("repos/%s/deployments/%d", namespace, d.Id))
		d.StatusesUrl = fmt.Sprintf("%s/statuses", d.Url)

		ds = append(ds, *d)
	}
	return ds, err
}

// Status represents the current status of a deployment
type Status struct {
	Url           string    `json:"url"`
	Id            int       `json:"id"`
	State         string    `json:"state"` // Required always
	Creator       User      `json:"creator"`
	Description   string    `json:"description"`
	TargetUrl     string    `json:"target_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeploymentUrl string    `json:"deployment_url"`
	RepositoryUrl string    `json:"repository_url"`
	DeploymentId  int       `json:"deploy_id"`
	Namespace     string    `json:"-"`
}
