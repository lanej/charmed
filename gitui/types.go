// Package gitui provides Git presentation components without repository access.
package gitui

import "time"

type RefKind string

const (
	LocalBranch  RefKind = "branch"
	RemoteBranch RefKind = "remote"
	Tag          RefKind = "tag"
)

type Ref struct {
	Name   string
	Kind   RefKind
	Target string
}

type HeadKind string

const (
	Attached HeadKind = "attached"
	Detached HeadKind = "detached"
	Unborn   HeadKind = "unborn"
)

type Head struct {
	Kind     HeadKind
	Ref      string
	CommitID string
}

type Commit struct {
	ID        string
	ParentIDs []string
	Subject   string
	Author    string
	Time      time.Time
}

// History requires child-before-parent order. Missing parents require Truncated.
type History struct {
	Commits   []Commit
	Refs      []Ref
	Head      Head
	Truncated bool
}

type ChangeKind string

const (
	Unchanged   ChangeKind = ""
	Added       ChangeKind = "added"
	Modified    ChangeKind = "modified"
	Deleted     ChangeKind = "deleted"
	Renamed     ChangeKind = "renamed"
	Copied      ChangeKind = "copied"
	TypeChanged ChangeKind = "type changed"
	Untracked   ChangeKind = "untracked"
	Conflicted  ChangeKind = "conflicted"
)

type ChangeSection string

const (
	Staged         ChangeSection = "staged"
	Unstaged       ChangeSection = "unstaged"
	UntrackedFiles ChangeSection = "untracked"
	Conflicts      ChangeSection = "conflicts"
)

type FileChange struct {
	ID           string
	Path         string
	OriginalPath string
	Index        ChangeKind
	Worktree     ChangeKind
	Conflict     string
}

// Known distinguishes a measured zero from an unavailable comparison.
type Divergence struct {
	Known  bool
	Ahead  int
	Behind int
}

// Repository.Unavailable distinguishes an unavailable observation from an empty working tree.
type Repository struct {
	Unavailable bool
	Reason      string
	Name        string
	Head        Head
	Upstream    string
	Divergence  Divergence
	Files       []FileChange
}

type Branch struct {
	Ref           Ref
	Current       bool
	Upstream      string
	Divergence    Divergence
	WorktreePaths []string
}

type DiffState string

const (
	TextDiff        DiffState = "text"
	BinaryDiff      DiffState = "binary"
	UnavailableDiff DiffState = "unavailable"
)

type LineKind string

const (
	ContextLine LineKind = "context"
	AddedLine   LineKind = "added"
	DeletedLine LineKind = "deleted"
)

type DiffLine struct {
	Kind      LineKind
	OldLine   int
	NewLine   int
	Text      string
	NoNewline bool
}

type Hunk struct {
	Header string
	Lines  []DiffLine
}

type FileDiff struct {
	FileID       string
	Section      ChangeSection
	Path         string
	OriginalPath string
	State        DiffState
	Reason       string
	Hunks        []Hunk
}

type DiffStat struct {
	State   DiffState
	Added   int
	Deleted int
}

type Stash struct {
	ID       string
	Ordinal  int
	Subject  string
	Branch   string
	CommitID string
	Diffs    []FileDiff
}

type Worktree struct {
	ID          string
	Path        string
	Head        Head
	Main        bool
	Current     bool
	Locked      bool
	LockReason  string
	Prunable    bool
	PruneReason string
	StatusKnown bool
	Status      Repository
}
