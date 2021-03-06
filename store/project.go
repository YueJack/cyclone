/*
Copyright 2017 caicloud authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package store

import (
	"fmt"
	"time"

	"github.com/caicloud/cyclone/pkg/api"
	"gopkg.in/mgo.v2/bson"
)

// CreateProject creates the project, returns the project created.
func (d *DataStore) CreateProject(project *api.Project) (*api.Project, error) {
	project.ID = bson.NewObjectId().Hex()
	project.CreatedTime = time.Now()
	project.UpdatedTime = time.Now()
	err := d.projectCollection.Insert(project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

// FindProjectByName finds the project by name. If find no project or more than one project, return error.
func (d *DataStore) FindProjectByName(name string) (*api.Project, error) {
	query := bson.M{"name": name}
	count, err := d.projectCollection.Find(query).Count()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, fmt.Errorf("there is no project with name %s", name)
	} else if count > 1 {
		return nil, fmt.Errorf("there are %d projects with the same name %s", count, name)
	}

	project := &api.Project{}
	err = d.projectCollection.Find(query).One(project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

// FindProjectByID finds the project by id.
func (d *DataStore) FindProjectByID(projectID string) (*api.Project, error) {
	project := &api.Project{}
	err := d.projectCollection.FindId(projectID).One(project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

// UpdateProject updates the project.
func (d *DataStore) UpdateProject(project *api.Project) error {
	updatedProject := *project
	updatedProject.UpdatedTime = time.Now()

	return d.projectCollection.UpdateId(project.ID, updatedProject)
}

// DeleteProjectByID deletes the project by id.
func (d *DataStore) DeleteProjectByID(projectID string) error {
	return d.projectCollection.RemoveId(projectID)
}
