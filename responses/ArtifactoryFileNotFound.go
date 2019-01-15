package responses

type ArtifactoryErrors struct  {
	Errors []error `json:"errors"`
}

type error struct  {
	Status int `json:"status"`
	Message string `json:"message"`
}