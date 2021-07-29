package yamltohtml_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/TutorialEdge/go-testing-bible/yamltohtml"
)

type TestCase struct {
	desc     string
	path     string
	expected string
}

func TestMain(m *testing.M) {
	fmt.Println("Hello World")
	ret := m.Run()
	fmt.Println("Tests have executed")
	os.Exit(ret)
}

func TestYamlToHTML(t *testing.T) {
	testCases := []TestCase{
		{
			desc:     "Test Case 1",
			path:     "testdata/test_01.yml",
			expected: "<html><head><title>My Awesome Page</title></head><body>This is my awesome content</body></html>",
		},
		{
			desc:     "Test Case 2",
			path:     "testdata/test_02.yml",
			expected: "<html><head><title>My Second Page</title></head><body>This is my awesome content</body></html>",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			result, err := yamltohtml.YamltoHTML(test.path)
			if err != nil {
				t.Fail()
			}

			t.Log(result)

			if result != test.expected {
				t.Fail()
			}
		})
	}

}
