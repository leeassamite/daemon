package models

import (
	"github.com/cSploit/daemon/models/internal"
	"github.com/ianschenck/envflag"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateProcessJob(t *testing.T) {
	envflag.Parse()
	internal.OpenDbForTests()

	pj, err := CreateProcessJob("date")

	require.Nil(t, err)
	require.NotNil(t, pj)
	require.Contains(t, commands, pj.JobId)
}

func TestFindProcessJob(t *testing.T) {
	envflag.Parse()
	internal.OpenDbForTests()

	pj, _ := CreateProcessJob("date")

	pj1, err := FindProcessJob(pj.JobId)

	require.Nil(t, err)
	require.NotNil(t, pj1)

	pj2, err := FindProcessJob(0)

	require.Nil(t, pj2)
	require.Error(t, err)
}

func TestProcessOutput(t *testing.T) {
	envflag.Parse()
	internal.OpenDbForTests()

	pj, _ := CreateProcessJob("date")

	cmd := commands[pj.JobId]

	require.NotNil(t, cmd)

	<-completed[pj.JobId]

	pj, _ = FindProcessJob(pj.JobId)

	require.NotNil(t, pj.ExitStatus)
	require.Equal(t, 0, *pj.ExitStatus)
	require.NotEmpty(t, pj.Output)
}
