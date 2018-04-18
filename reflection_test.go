package crazytalk

import (
	"fmt"
	"net"
	"os"
	"testing"

	"github.com/daveamit/crazytalk/test"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var ct CrazyTalk

func TestMain(m *testing.M) {
	// Setup
	code := 1
	defer func() {
		p := recover()
		if p != nil {
			fmt.Fprintf(os.Stderr, "PANIC: %v\n", p)
		}
		os.Exit(code)
	}()

	// Start up a server on an ephemeral port
	l, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		panic(fmt.Sprintf("Failed to listen to port: %s", err.Error()))
	}
	svr := grpc.NewServer()
	test.Register(svr)
	reflection.Register(svr)

	go svr.Serve(l)
	defer svr.Stop()

	ct = NewReflectionCrazyTalk("127.0.0.1:3000")

	// Execute
	code = m.Run()

	// Tear-Down
	// Tear down done using defer statements
}

func TestListService(t *testing.T) {

	list, err := ct.ListServices()
	if err != nil {
		t.Error(err)
	}

	expected := []Service{
		// Service{
		// 	Name: "ServerReflection",
		// 	Methods: []Method{
		// 		Method{
		// 			Name: "ServerReflectionInfo",
		// 			InputType: Type{
		// 				Name: "ServerReflectionRequest",
		// 				Fields: []Field{
		// 					Field{
		// 						Name: "host",
		// 					},
		// 					Field{
		// 						Name: "file_by_filename",
		// 					},
		// 					Field{
		// 						Name: "file_containing_symbol",
		// 					},
		// 					Field{
		// 						Name: "file_containing_extension",
		// 					},
		// 					Field{
		// 						Name: "all_extension_numbers_of_type",
		// 					},
		// 					Field{
		// 						Name: "list_services",
		// 					},
		// 				},
		// 			},
		// 			OutputType: Type{
		// 				Name: "ServerReflectionResponse",
		// 				Fields: []Field{
		// 					Field{
		// 						Name: "valid_host",
		// 					},
		// 					Field{
		// 						Name: "original_request",
		// 					},
		// 					Field{
		// 						Name: "file_descriptor_response",
		// 					},
		// 					Field{
		// 						Name: "all_extension_numbers_response",
		// 					},
		// 					Field{
		// 						Name: "list_services_response",
		// 					},
		// 					Field{
		// 						Name: "error_response",
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// },
		Service{
			Name:               "Hello",
			FullyQualifiedName: "test.Hello",
			Methods: []Method{
				Method{
					Name:               "SayHi",
					FullyQualifiedName: "test.Hello.SayHi",
					InputType: Type{
						Name:               "SayHiRequest",
						FullyQualifiedName: "test.SayHiRequest",
						Fields: []Field{
							Field{
								Name:       "name",
								ActualType: "TYPE_STRING",
							},
							Field{
								Name:       "p",
								ActualType: "TYPE_MESSAGE",
								Type: Type{
									Name:               "P",
									FullyQualifiedName: "test.P",
									Fields: []Field{
										Field{
											Name:       "cFromP",
											ActualType: "TYPE_MESSAGE",
											Type: Type{
												Name:               "C",
												FullyQualifiedName: "test.C",
												Fields: []Field{
													Field{
														Name:       "pFromC",
														ActualType: "TYPE_MESSAGE",
														Type: Type{
															Name:                    "P",
															FullyQualifiedName:      "test.P",
															TruncatedDueToRecursion: true,
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					OutputType: Type{
						Name:               "SayHiResponse",
						FullyQualifiedName: "test.SayHiResponse",
						Fields: []Field{
							Field{
								Name:       "message",
								ActualType: "TYPE_STRING",
							},
						},
					},
				},
			},
		},
	}

	assert.EqualValues(t, expected, list)
}

func TestInvokeRPC(t *testing.T) {

	ct.ListServices()
	response, err := ct.InvokeRPC("test.Hello.SayHi", "{\"name\": \"Dave\"}")
	if err != nil {
		t.Error(err)
	}

	assert.EqualValues(t, "{\"message\":\"Hello Dave!\"}", response)
}
