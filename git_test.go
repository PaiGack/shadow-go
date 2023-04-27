package shadow_go

import "testing"

func TestGit_CurrentBranch(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gitCommand{}
			branch, err := g.currentBranch()
			if err != nil {
				t.Fatal(err)
			}
			t.Log(branch)
		})
	}
}

func TestGitCommand_Clean(t *testing.T) {
	tests := []struct {
		name    string
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gitCommand{}
			got, err := g.clean()
			if (err != nil) != tt.wantErr {
				t.Errorf("clean() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("clean() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitCommand_LastTag(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gitCommand{}
			got, err := g.lastTag()
			if (err != nil) != tt.wantErr {
				t.Errorf("lastTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("lastTag() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitCommand_CurrentTag(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gitCommand{}
			got, err := g.currentTag()
			if (err != nil) != tt.wantErr {
				t.Errorf("currentTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("currentTag() got = %v, want %v", got, tt.want)
			}
		})
	}
}
