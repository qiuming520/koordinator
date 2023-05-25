/*
Copyright 2022 The Koordinator Authors.

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

package metriccache

var (
	defaultMetricFactory = NewMetricFactory()

	// define all kinds of MetricResource
	NodeCPUUsageMetric          = defaultMetricFactory.New(NodeMetricCPUUsage)
	NodeMemoryUsageMetric       = defaultMetricFactory.New(NodeMetricMemoryUsage)
	NodeGPUCoreUsageMetric      = defaultMetricFactory.New(NodeMetricGPUCoreUsage).withPropertySchema(MetricPropertyGPUMinor, MetricPropertyGPUDeviceUUID)
	NodeGPUMemUsageMetric       = defaultMetricFactory.New(NodeMetricGPUMemUsage).withPropertySchema(MetricPropertyGPUMinor, MetricPropertyGPUDeviceUUID)
	NodeGPUTotalMetric          = defaultMetricFactory.New(NodeMetricGPUMemTotal).withPropertySchema(MetricPropertyGPUMinor, MetricPropertyGPUDeviceUUID)
	PodCPUUsageMetric           = defaultMetricFactory.New(PodMetricCPUUsage).withPropertySchema(MetricPropertyPodUID)
	PodCPUThrottledMetric       = defaultMetricFactory.New(PodMetricCPUThrottled).withPropertySchema(MetricPropertyPodUID)
	ContainerCPUUsageMetric     = defaultMetricFactory.New(ContainerMetricCPUUsage).withPropertySchema(MetricPropertyContainerID)
	ContainerCPUThrottledMetric = defaultMetricFactory.New(ContainerMetricCPUThrottled).withPropertySchema(MetricPropertyContainerID)
)
