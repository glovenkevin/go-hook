package request

type DeployHookRequest struct {
	RepoPath string `json:"repo_path" validate:"required"`
}
