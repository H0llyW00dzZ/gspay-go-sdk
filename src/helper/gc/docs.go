// Copyright 2026 H0llyW00dzZ
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package gc provides buffer pool management for efficient memory reuse.
//
// This package wraps bytebufferpool to provide a consistent interface for
// buffer pooling, reducing memory allocations in high-throughput scenarios
// such as API request/response handling.
//
// # Core Types
//
// The package defines two main interfaces:
//
//   - [Buffer]: A reusable byte buffer implementing io.Writer, io.WriterTo,
//     and io.ReaderFrom interfaces
//   - [Pool]: A buffer pool for acquiring and releasing buffers
//
// # Default Pool
//
// The [Default] pool is the recommended way to use buffer pooling:
//
//	buf := gc.Default.Get()
//	defer func() {
//	    buf.Reset()
//	    gc.Default.Put(buf)
//	}()
//
//	// Use the buffer for I/O operations
//	buf.WriteString("Hello, World!")
//	data := buf.String()
//
// # Reading HTTP Response
//
//	buf := gc.Default.Get()
//	defer func() {
//	    buf.Reset()
//	    gc.Default.Put(buf)
//	}()
//
//	if _, err := buf.ReadFrom(resp.Body); err != nil {
//	    return err
//	}
//
//	var result Response
//	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
//	    return err
//	}
//
// # Writing HTTP Request
//
//	buf := gc.Default.Get()
//	defer func() {
//	    buf.Reset()
//	    gc.Default.Put(buf)
//	}()
//
//	if err := json.NewEncoder(buf).Encode(request); err != nil {
//	    return err
//	}
//
//	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, buf)
//
// # Thread Safety
//
// The [Default] pool and all [Pool] implementations are safe for concurrent
// use by multiple goroutines.
//
// # Memory Efficiency
//
// Buffer pooling provides significant memory savings in high-concurrency
// environments by reusing allocations instead of constant allocation and
// garbage collection. This is especially beneficial for payment processing
// where many API requests are handled simultaneously.
package gc
