/*
 * Copyright (c) 2024. Devtron Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package chartRepo

import (
	"context"
	"github.com/devtron-labs/common-lib/utils/k8s"
	chartRepoRepository "github.com/devtron-labs/devtron/pkg/chartRepo/repository"
	cluster2 "github.com/devtron-labs/devtron/pkg/cluster"
	"github.com/devtron-labs/devtron/pkg/cluster/repository"
	"github.com/go-pg/pg"
	"github.com/stretchr/testify/mock"
)

type ChartRepoRepositoryImplMock struct {
	mock.Mock
}

func (impl ChartRepoRepositoryImplMock) Save(chartRepo *chartRepoRepository.ChartRepo, tx *pg.Tx) error {
	panic("implement me")
}
func (impl ChartRepoRepositoryImplMock) Update(chartRepo *chartRepoRepository.ChartRepo, tx *pg.Tx) error {
	panic("implement me")
}
func (impl ChartRepoRepositoryImplMock) GetDefault() (*chartRepoRepository.ChartRepo, error) {
	panic("implement me")
}
func (impl ChartRepoRepositoryImplMock) FindById(id int) (*chartRepoRepository.ChartRepo, error) {
	panic("implement me")
}
func (impl ChartRepoRepositoryImplMock) FindAll() ([]*chartRepoRepository.ChartRepo, error) {
	panic("implement me")
}
func (impl ChartRepoRepositoryImplMock) GetConnection() *pg.DB {
	panic("implement me")
}
func (impl ChartRepoRepositoryImplMock) MarkChartRepoDeleted(chartRepo *chartRepoRepository.ChartRepo, tx *pg.Tx) error {
	panic("implement me")
}

// ----------
type ClusterServiceImplMock struct {
	mock.Mock
}

//func (impl ClusterServiceImplMock) Save(chartRepo *chartRepoRepository.ChartRepo, tx *pg.Tx) error {
//	panic("implement me")
//}

func (impl ClusterServiceImplMock) Save(parent context.Context, bean *cluster2.ClusterBean, userId int32) (*cluster2.ClusterBean, error) {
	panic("implement me")
}
func (impl ClusterServiceImplMock) FindOne(clusterName string) (*cluster2.ClusterBean, error) {
	panic("implement me")
}
func (impl ClusterServiceImplMock) FindOneActive(clusterName string) (*cluster2.ClusterBean, error) {
	panic("implement me")
}
func (impl ClusterServiceImplMock) FindAll() ([]*cluster2.ClusterBean, error) {
	panic("implement me")
}
func (impl ClusterServiceImplMock) FindAllActive() ([]cluster2.ClusterBean, error) {
	panic("implement me")
}
func (impl ClusterServiceImplMock) DeleteFromDb(bean *cluster2.ClusterBean, userId int32) error {
	panic("implement me")
}

func (impl ClusterServiceImplMock) FindById(id int) (*cluster2.ClusterBean, error) {
	panic("implement me")
}
func (impl ClusterServiceImplMock) FindByIds(id []int) ([]cluster2.ClusterBean, error) {
	panic("implement me")
}
func (impl ClusterServiceImplMock) Update(ctx context.Context, bean *cluster2.ClusterBean, userId int32) (*cluster2.ClusterBean, error) {
	panic("implement me")
}
func (impl ClusterServiceImplMock) Delete(bean *cluster2.ClusterBean, userId int32) error {
	panic("implement me")
}

func (impl ClusterServiceImplMock) FindAllForAutoComplete() ([]cluster2.ClusterBean, error) {
	panic("implement me")
}
func (impl ClusterServiceImplMock) CreateGrafanaDataSource(clusterBean *cluster2.ClusterBean, env *repository.Environment) (int, error) {
	panic("implement me")
}
func (impl ClusterServiceImplMock) GetClusterConfig(cluster *cluster2.ClusterBean) (*k8s.ClusterConfig, error) {
	panic("implement me")
}
