/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type SaveOptions struct {

	//  The type of save operation to be performed.
	SaveType string `json:"SaveType,omitempty"`

	//  The path to the file that will container the saved state.
	SaveStateFilePath string `json:"SaveStateFilePath,omitempty"`
}
