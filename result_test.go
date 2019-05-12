package cliassert

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestResult_Show(t *testing.T) {
	cases := []struct {
		Name   string
		Result Result
	}{
		{"zero", Result{}},
		{"successes", Result{
			successes: []string{"ok1", "ok2"},
		}},
		{"failures", Result{
			failures: []string{"ng1", "ng2"},
		}},
		{"combined", Result{
			failures:  []string{"ng1", "ng2"},
			successes: []string{"ok1", "ok2"},
		}},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got, err := c.Result.Show()
			if err != nil {
				t.Fatalf("Show return error: %v", err)
			}
			golden := filepath.Join("fixtures/result", "show_"+c.Name+".golden")
			if *update {
				ioutil.WriteFile(golden, []byte(got), 0644)
			}
			data, _ := ioutil.ReadFile(golden)
			want := string(data)
			if got != want {
				t.Errorf("Show \ngot:\n%v,want:\n%v", got, want)
			}
		})
	}
}

func TestResult_ShowDetails(t *testing.T) {
	cases := []struct {
		Name   string
		Result Result
	}{
		{"zero", Result{}},
		{"successes", Result{
			exitStatus: "1",
			stdout:     "stdout\n",
			stderr:     "stderr\n",
			successes:  []string{"ok1", "ok2"},
		}},
		{"failures", Result{
			exitStatus: "1",
			stdout:     "stdout\n",
			stderr:     "stderr\n",
			failures:   []string{"ng1", "ng2"},
		}},
		{"combined", Result{
			exitStatus: "1",
			stdout:     "stdout\n",
			stderr:     "stderr\n",
			failures:   []string{"ng1", "ng2"},
			successes:  []string{"ok1", "ok2"},
		}},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got, err := c.Result.ShowDetails()
			if err != nil {
				t.Fatalf("Show return error: %v", err)
			}
			golden := filepath.Join("fixtures/result", "show_details_"+c.Name+".golden")
			if *update {
				ioutil.WriteFile(golden, []byte(got), 0644)
			}
			data, _ := ioutil.ReadFile(golden)
			want := string(data)
			if got != want {
				t.Errorf("Show \ngot:\n%v,want:\n%v", got, want)
			}
		})
	}
}

func TestResult_Stdout(t *testing.T) {
	cases := []struct {
		Name string
		Want string
	}{
		{"ok", "ok"},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			r := Result{stdout: c.Want}
			got := r.Stdout()
			if got != c.Want {
				t.Errorf("Stdout got:%v, want:%v", got, c.Want)
			}
		})
	}
}
