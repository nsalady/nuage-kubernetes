/*
  Copyright (c) 2015, Alcatel-Lucent Inc
  All rights reserved.

  Redistribution and use in source and binary forms, with or without
  modification, are permitted provided that the following conditions are met:
      * Redistributions of source code must retain the above copyright
        notice, this list of conditions and the following disclaimer.
      * Redistributions in binary form must reproduce the above copyright
        notice, this list of conditions and the following disclaimer in the
        documentation and/or other materials provided with the distribution.
      * Neither the name of the copyright holder nor the names of its contributors
        may be used to endorse or promote products derived from this software without
        specific prior written permission.

  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
  ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
  WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY
  DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
  (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
  LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
  ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
  (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
  SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// PublicNetworkMacroIdentity represents the Identity of the object
var PublicNetworkMacroIdentity = bambou.Identity{
	Name:     "publicnetwork",
	Category: "publicnetworks",
}

// PublicNetworkMacrosList represents a list of PublicNetworkMacros
type PublicNetworkMacrosList []*PublicNetworkMacro

// PublicNetworkMacrosAncestor is the interface of an ancestor of a PublicNetworkMacro must implement.
type PublicNetworkMacrosAncestor interface {
	PublicNetworkMacros(*bambou.FetchingInfo) (PublicNetworkMacrosList, *bambou.Error)
	CreatePublicNetworkMacros(*PublicNetworkMacro) *bambou.Error
}

// PublicNetworkMacro represents the model of a publicnetwork
type PublicNetworkMacro struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	IPType        string `json:"IPType,omitempty"`
	IPv6Address   string `json:"IPv6Address,omitempty"`
	Name          string `json:"name,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	Address       string `json:"address,omitempty"`
	Netmask       string `json:"netmask,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewPublicNetworkMacro returns a new *PublicNetworkMacro
func NewPublicNetworkMacro() *PublicNetworkMacro {

	return &PublicNetworkMacro{}
}

// Identity returns the Identity of the object.
func (o *PublicNetworkMacro) Identity() bambou.Identity {

	return PublicNetworkMacroIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PublicNetworkMacro) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PublicNetworkMacro) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PublicNetworkMacro from the server
func (o *PublicNetworkMacro) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PublicNetworkMacro into the server
func (o *PublicNetworkMacro) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PublicNetworkMacro from the server
func (o *PublicNetworkMacro) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the PublicNetworkMacro
func (o *PublicNetworkMacro) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the PublicNetworkMacro
func (o *PublicNetworkMacro) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the PublicNetworkMacro
func (o *PublicNetworkMacro) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the PublicNetworkMacro
func (o *PublicNetworkMacro) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the PublicNetworkMacro
func (o *PublicNetworkMacro) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}

// CreateEventLog creates a new child EventLog under the PublicNetworkMacro
func (o *PublicNetworkMacro) CreateEventLog(child *EventLog) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
