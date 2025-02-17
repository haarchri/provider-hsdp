/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type UserObservation struct {

	// The GUID of the user
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type UserParameters struct {

	// (Semi-Required) The email address of the user
	// The email address of the user.
	// +kubebuilder:validation:Required
	Email *string `json:"email" tf:"email,omitempty"`

	// First name of the user
	// The first name of the user.
	// +kubebuilder:validation:Required
	FirstName *string `json:"firstName" tf:"first_name,omitempty"`

	// Last name of the user
	// The last name of the user.
	// +kubebuilder:validation:Required
	LastName *string `json:"lastName" tf:"last_name,omitempty"`

	// The login ID of the user (NEW since v0.4.0)
	// +kubebuilder:validation:Required
	Login *string `json:"login" tf:"login,omitempty"`

	// Mobile number of the user. E.164 format
	// The optional mobile phone number of the user.
	// +kubebuilder:validation:Optional
	Mobile *string `json:"mobile,omitempty" tf:"mobile,omitempty"`

	// The managing organization of the user
	// The managing organization of the user.
	// +crossplane:generate:reference:type=Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Selector for a Organization to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationIDSelector *v1.Selector `json:"organizationIdSelector,omitempty" tf:"-"`

	// Reference to a Organization to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// When specified this will skip the email activation
	// flow and immediately activate the IAM account. Very Important: you are responsible
	// for sharing this password with the new IAM user through some channel of communication.
	// No email will be triggered by the system. If unsure, do not set a password so the normal
	// email activation flow is followed. Finally, any password value changes after user creation
	// will have no effect on the users' actual password.
	// When specified this will skip the email activation flow and immediately activate the IAM account. Very Important: you are responsible for sharing this password with the new IAM user through some channel of communication. No email will be triggered by the system. If unsure, do not set a password so the normal email activation flow is followed. Finally, any password value changes after user creation will have no effect on the users' actual password.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Preferred communication channel.
	// Email and SMS are supported channels. Email is the default channel if e-mail address is provided.
	// Values supported: [ email | sms ]
	// Preferred communication channel. Email and SMS are supported channels. Email is the default channel if e-mail address is provided. Values supported: [ email | sms ].
	// +kubebuilder:validation:Optional
	PreferredCommunicationChannel *string `json:"preferredCommunicationChannel,omitempty" tf:"preferred_communication_channel,omitempty"`

	// Language preference for all communications.
	// Value can be a two letter language code as defined by ISO 639-1 (en, de) or it can be a combination
	// of language code and country code (en-gb, en-us). The country code is as per ISO 3166 two letter code (alpha-2)
	// Language preference for all communications. Value can be a two letter language code as defined by ISO 639-1 (en, de) or it can be a combination of language code and country code (en-gb, en-us). The country code is as per ISO 3166 two letter code (alpha-2).
	// +kubebuilder:validation:Optional
	PreferredLanguage *string `json:"preferredLanguage,omitempty" tf:"preferred_language,omitempty"`

	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// User is the Schema for the Users API. Manages HSDP IAM Users
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hsdp}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserSpec   `json:"spec"`
	Status            UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
