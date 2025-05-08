package api

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/kcchandra/golang-hook/api/request"
	httpext "github.com/kcchandra/golang-hook/pkg/http"
)

func DeployHook(w http.ResponseWriter, r *http.Request) {
	var req request.DeployHookRequest

	if err := httpext.ReadJSON(r, &req); err != nil {
		httpext.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Execute git pull and pm2 reload
	cmd := exec.Command("sh", "-c", "git pull && pm2 reload ecosystem.config.js")
	cmd.Dir = req.RepoPath

	output, err := cmd.CombinedOutput()
	if err != nil {
		httpext.Error(w, http.StatusInternalServerError, "Failed to execute deploy commands: "+err.Error())
		return
	}

	// Log the output for debugging
	log.Printf("Deploy output: %s", string(output))

	httpext.StatusOK(w, "result: "+string(output))
}
