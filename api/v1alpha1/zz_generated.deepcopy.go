//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*


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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"Fluid-Datatable/pkg/common"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIGatewayStatus) DeepCopyInto(out *APIGatewayStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIGatewayStatus.
func (in *APIGatewayStatus) DeepCopy() *APIGatewayStatus {
	if in == nil {
		return nil
	}
	out := new(APIGatewayStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlluxioCompTemplateSpec) DeepCopyInto(out *AlluxioCompTemplateSpec) {
	*out = *in
	if in.JvmOptions != nil {
		in, out := &in.JvmOptions, &out.JvmOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make(map[string]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlluxioCompTemplateSpec.
func (in *AlluxioCompTemplateSpec) DeepCopy() *AlluxioCompTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(AlluxioCompTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlluxioFuseSpec) DeepCopyInto(out *AlluxioFuseSpec) {
	*out = *in
	if in.JvmOptions != nil {
		in, out := &in.JvmOptions, &out.JvmOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlluxioFuseSpec.
func (in *AlluxioFuseSpec) DeepCopy() *AlluxioFuseSpec {
	if in == nil {
		return nil
	}
	out := new(AlluxioFuseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlluxioRuntime) DeepCopyInto(out *AlluxioRuntime) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlluxioRuntime.
func (in *AlluxioRuntime) DeepCopy() *AlluxioRuntime {
	if in == nil {
		return nil
	}
	out := new(AlluxioRuntime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlluxioRuntime) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlluxioRuntimeList) DeepCopyInto(out *AlluxioRuntimeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlluxioRuntime, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlluxioRuntimeList.
func (in *AlluxioRuntimeList) DeepCopy() *AlluxioRuntimeList {
	if in == nil {
		return nil
	}
	out := new(AlluxioRuntimeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlluxioRuntimeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlluxioRuntimeSpec) DeepCopyInto(out *AlluxioRuntimeSpec) {
	*out = *in
	out.AlluxioVersion = in.AlluxioVersion
	in.Master.DeepCopyInto(&out.Master)
	in.JobMaster.DeepCopyInto(&out.JobMaster)
	in.Worker.DeepCopyInto(&out.Worker)
	in.JobWorker.DeepCopyInto(&out.JobWorker)
	in.APIGateway.DeepCopyInto(&out.APIGateway)
	in.InitUsers.DeepCopyInto(&out.InitUsers)
	in.Fuse.DeepCopyInto(&out.Fuse)
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.JvmOptions != nil {
		in, out := &in.JvmOptions, &out.JvmOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.TieredStore.DeepCopyInto(&out.TieredStore)
	out.Data = in.Data
	if in.RunAs != nil {
		in, out := &in.RunAs, &out.RunAs
		*out = new(User)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlluxioRuntimeSpec.
func (in *AlluxioRuntimeSpec) DeepCopy() *AlluxioRuntimeSpec {
	if in == nil {
		return nil
	}
	out := new(AlluxioRuntimeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in CacheStateList) DeepCopyInto(out *CacheStateList) {
	{
		in := &in
		*out = make(CacheStateList, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheStateList.
func (in CacheStateList) DeepCopy() CacheStateList {
	if in == nil {
		return nil
	}
	out := new(CacheStateList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheableNodeAffinity) DeepCopyInto(out *CacheableNodeAffinity) {
	*out = *in
	if in.Required != nil {
		in, out := &in.Required, &out.Required
		*out = new(v1.NodeSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheableNodeAffinity.
func (in *CacheableNodeAffinity) DeepCopy() *CacheableNodeAffinity {
	if in == nil {
		return nil
	}
	out := new(CacheableNodeAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Data) DeepCopyInto(out *Data) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Data.
func (in *Data) DeepCopy() *Data {
	if in == nil {
		return nil
	}
	out := new(Data)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataTable) DeepCopyInto(out *DataTable) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataTable.
func (in *DataTable) DeepCopy() *DataTable {
	if in == nil {
		return nil
	}
	out := new(DataTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataTable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataTableList) DeepCopyInto(out *DataTableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataTable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataTableList.
func (in *DataTableList) DeepCopy() *DataTableList {
	if in == nil {
		return nil
	}
	out := new(DataTableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataTableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataTableSpec) DeepCopyInto(out *DataTableSpec) {
	*out = *in
	if in.Schemas != nil {
		in, out := &in.Schemas, &out.Schemas
		*out = make([]Schema, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Runtimes != nil {
		in, out := &in.Runtimes, &out.Runtimes
		*out = make([]Runtime, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataTableSpec.
func (in *DataTableSpec) DeepCopy() *DataTableSpec {
	if in == nil {
		return nil
	}
	out := new(DataTableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataTableStatus) DeepCopyInto(out *DataTableStatus) {
	*out = *in
	if in.Schemas != nil {
		in, out := &in.Schemas, &out.Schemas
		*out = make([]Schema, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]DatatableCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CacheStates != nil {
		in, out := &in.CacheStates, &out.CacheStates
		*out = make(CacheStateList, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataTableStatus.
func (in *DataTableStatus) DeepCopy() *DataTableStatus {
	if in == nil {
		return nil
	}
	out := new(DataTableStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatatableCondition) DeepCopyInto(out *DatatableCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatatableCondition.
func (in *DatatableCondition) DeepCopy() *DatatableCondition {
	if in == nil {
		return nil
	}
	out := new(DatatableCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCFSStatus) DeepCopyInto(out *HCFSStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCFSStatus.
func (in *HCFSStatus) DeepCopy() *HCFSStatus {
	if in == nil {
		return nil
	}
	out := new(HCFSStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitUsersSpec) DeepCopyInto(out *InitUsersSpec) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitUsersSpec.
func (in *InitUsersSpec) DeepCopy() *InitUsersSpec {
	if in == nil {
		return nil
	}
	out := new(InitUsersSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Level) DeepCopyInto(out *Level) {
	*out = *in
	if in.Quota != nil {
		in, out := &in.Quota, &out.Quota
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Level.
func (in *Level) DeepCopy() *Level {
	if in == nil {
		return nil
	}
	out := new(Level)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Runtime) DeepCopyInto(out *Runtime) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Runtime.
func (in *Runtime) DeepCopy() *Runtime {
	if in == nil {
		return nil
	}
	out := new(Runtime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeCondition) DeepCopyInto(out *RuntimeCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeCondition.
func (in *RuntimeCondition) DeepCopy() *RuntimeCondition {
	if in == nil {
		return nil
	}
	out := new(RuntimeCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeStatus) DeepCopyInto(out *RuntimeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]RuntimeCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CacheStates != nil {
		in, out := &in.CacheStates, &out.CacheStates
		*out = make(common.CacheStateList, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.APIGatewayStatus != nil {
		in, out := &in.APIGatewayStatus, &out.APIGatewayStatus
		*out = new(APIGatewayStatus)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeStatus.
func (in *RuntimeStatus) DeepCopy() *RuntimeStatus {
	if in == nil {
		return nil
	}
	out := new(RuntimeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Schema) DeepCopyInto(out *Schema) {
	*out = *in
	if in.Tables != nil {
		in, out := &in.Tables, &out.Tables
		*out = make([]Table, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Schema.
func (in *Schema) DeepCopy() *Schema {
	if in == nil {
		return nil
	}
	out := new(Schema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Table) DeepCopyInto(out *Table) {
	*out = *in
	if in.ColumnName != nil {
		in, out := &in.ColumnName, &out.ColumnName
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PartitionColumn != nil {
		in, out := &in.PartitionColumn, &out.PartitionColumn
		*out = make([]map[string]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Table.
func (in *Table) DeepCopy() *Table {
	if in == nil {
		return nil
	}
	out := new(Table)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TieredStore) DeepCopyInto(out *TieredStore) {
	*out = *in
	if in.Levels != nil {
		in, out := &in.Levels, &out.Levels
		*out = make([]Level, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TieredStore.
func (in *TieredStore) DeepCopy() *TieredStore {
	if in == nil {
		return nil
	}
	out := new(TieredStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
	if in.UID != nil {
		in, out := &in.UID, &out.UID
		*out = new(int64)
		**out = **in
	}
	if in.GID != nil {
		in, out := &in.GID, &out.GID
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionSpec) DeepCopyInto(out *VersionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionSpec.
func (in *VersionSpec) DeepCopy() *VersionSpec {
	if in == nil {
		return nil
	}
	out := new(VersionSpec)
	in.DeepCopyInto(out)
	return out
}