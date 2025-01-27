//go:build !no_runtime_type_checking

package aspects

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_SqsRequireSsl) validateVisitParameters(node constructs.IConstruct) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

