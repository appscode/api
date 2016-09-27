package kubernetes_v02

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	patterns := make([]runtime.Pattern, 0)

	patterns = append(patterns, pattern_Clients_Copy_0)
	patterns = append(patterns, pattern_Clients_Delete_0)
	patterns = append(patterns, pattern_Clients_Describe_0)
	patterns = append(patterns, pattern_Clients_List_0)
	patterns = append(patterns, pattern_Clients_Update_0)
	patterns = append(patterns, pattern_Clients_UpdateConfigMap_0)
	patterns = append(patterns, pattern_Clients_UpdateSecret_0)

	return patterns
}
