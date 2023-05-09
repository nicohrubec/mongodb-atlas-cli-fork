// Copyright 2021 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package convert

import (
	atlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
	"strings"

	"go.mongodb.org/ops-manager/opsmngr"
)

const (
	AdminDB             = "admin"
	ExternalAuthDB      = "$external"
	roleSep             = "@"
	scopeSep            = ":"
	collectionSep       = "."
	defaultUserDatabase = "admin"
	defaultResourceType = "CLUSTER"
)

// BuildAtlasRoles converts the roles inside the array of string in an array of mongodbatlas.Role structs.
// r contains roles in the format roleName@dbName.
func BuildAtlasRoles(r []string) []atlasv2.Role {
	roles := make([]atlasv2.Role, len(r))
	for i, roleP := range r {
		roleName, databaseName := splitRoleAndDBName(roleP)
		dbCollection := strings.Split(databaseName, collectionSep)
		databaseName = dbCollection[0]
		roles[i] = atlasv2.Role{
			RoleName:       roleName,
			DatabaseName:   databaseName,
			CollectionName: buildCollectionName(dbCollection),
		}
	}
	return roles
}

func buildCollectionName(dbCollection []string) *string {
	var collectionName string
	if len(dbCollection) > 1 {
		collectionName = strings.Join(dbCollection[1:], ".")
	}
	return getStringPointerIfNotEmpty(collectionName)
}

// GetAuthDB determines the authentication database based on the type of user.
// LDAP, X509 and AWSIAM should all use $external.
// SCRAM-SHA should use admin.
func GetAuthDB(user *atlasv2.DatabaseUser)  (name string) {
	// base documentation https://registry.terraform.io/providers/mongodb/mongodbatlas/latest/docs/resources/database_user
	name = "admin"
	_, isX509 := adminX509Type[getStringIfPointerNotNil(user.X509Type)]
	_, isIAM := awsIAMType[getStringIfPointerNotNil(user.AwsIAMType)]

	// just USER is external
	isLDAP := user.LdapAuthType != nil && len(*user.LdapAuthType) > 0 && *user.LdapAuthType == "USER"

	if isX509 || isIAM || isLDAP {
		return "$external"
	}

	return name
}

var adminX509Type = map[string]struct{}{
	"MANAGED":  {},
	"CUSTOMER": {},
}

var awsIAMType = map[string]struct{}{
	"USER": {},
	"ROLE": {},
}

func getStringPointerIfNotEmpty(input string) *string{
	if input != "" {
		return &input
	}
	return nil
}

func getStringIfPointerNotNil(input *string) string{
	if input != nil {
		return *input
	}
	return ""
}

func splitRoleAndDBName(roleAndDBNAme string) (role, dbName string) {
	rd := strings.Split(roleAndDBNAme, roleSep)
	dbName = defaultUserDatabase
	role = rd[0]
	if len(rd) > 1 {
		dbName = rd[1]
	}
	return
}

// BuildOMRoles converts the roles inside the array of string in an array of opsmngr.Role structs.
// r contains roles in the format roleName@dbName.
func BuildOMRoles(r []string) []*opsmngr.Role {
	roles := make([]*opsmngr.Role, len(r))
	for i, roleP := range r {
		roleName, databaseName := splitRoleAndDBName(roleP)

		roles[i] = &opsmngr.Role{
			Role:     roleName,
			Database: databaseName,
		}
	}
	return roles
}

// BuildAtlasScopes converts the scopes inside the array of string in an array of mongodbatlas.Scope structs.
// r contains resources in the format resourceName:resourceType.
func BuildAtlasScopes(r []string) []atlasv2.UserScope {
	scopes := make([]atlasv2.UserScope, len(r))
	for i, scopeP := range r {
		scope := strings.Split(scopeP, scopeSep)
		resourceType := defaultResourceType
		if len(scope) > 1 {
			resourceType = scope[1]
		}

		scopes[i] = atlasv2.UserScope{
			Name: scope[0],
			Type: strings.ToUpper(resourceType),
		}
	}
	return scopes
}
