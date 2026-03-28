package cli

import (
	"apix/internal/types"
	"fmt"
	"strings"
)

// type ResponseParams struct {
// 	Method      string
// 	Endpoint    string
// 	Status      int
// 	ContentType string
// 	Body        string
// 	Duration    int64
// 	Size        string
// 	Headers     map[string][]string
// 	Timing      *Timing
// }

// type Timing struct {
// 	DNS  int64
// 	TCP  int64
// 	TLS  int64
// 	TTFB int64
// }

// func RenderResponse(params ResponseParams) {

// 	// 🔹 Top line (compact summary)
// 	fmt.Printf("%s %s  →  %s\n\n",
// 		Cyan(params.Method),
// 		params.Endpoint,
// 		FormatStatus(params.Status),
// 	)

// 	// 🔹 Headers (only useful ones or all if you prefer)
// 	if len(params.Headers) > 0 {
// 		fmt.Println(Muted("Headers"))

// 		for k, v := range params.Headers {
// 			fmt.Printf("%s: %s\n",
// 				strings.ToLower(k),
// 				strings.Join(v, ", "),
// 			)
// 		}

// 		fmt.Println()
// 	}

// 	// 🔹 Body
// 	fmt.Println(Muted("Body"))

// 	if strings.Contains(params.ContentType, "application/json") {
// 		formatted, err := PrettyJSON([]byte(params.Body))
// 		if err != nil {
// 			fmt.Println(params.Body)
// 		} else {
// 			fmt.Println(formatted)
// 		}
// 	} else {
// 		fmt.Println(params.Body)
// 	}

// 	fmt.Println()

//		// 🔹 Footer (compact)
//		fmt.Println(Muted(fmt.Sprintf("⏱ %d ms • %s", params.Duration, params.Size)))
//	}
func RenderResponse(params types.ResponseParams, verbose bool) {

	// ▌ REQUEST
	fmt.Println(Cyan("▌ REQUEST"))
	printKV("Method", params.Method)
	printKV("Endpoint", params.Endpoint)

	fmt.Println()

	// ▌ RESPONSE
	fmt.Println(Cyan("▌ RESPONSE"))
	printKV("Status", FormatStatus(params.Status))
	printKV("Duration", fmt.Sprintf("%d ms", params.Duration))
	printKV("Size", params.Size)

	if params.ContentType != "" {
		printKV("Type", params.ContentType)
	}

	fmt.Println()

	// ▌ HEADERS (only in verbose OR if you want always, tweak here)
	if verbose && len(params.Headers) > 0 {
		fmt.Println(Cyan("▌ HEADERS"))
		for k, v := range params.Headers {
			fmt.Printf("  %s: %s\n", strings.ToLower(k), strings.Join(v, ", "))
		}
		fmt.Println()
	}

	// ▌ BODY
	fmt.Println(Cyan("▌ BODY"))
	fmt.Println(formatBody(params.Body, params.ContentType))
	fmt.Println()

	// ▌ TIMING (only verbose)
	timing := params.Timing

	if verbose && timing != nil {
		fmt.Println(Cyan("▌ TIMING"))
		printKV("DNS", fmt.Sprintf("%d ms", timing.DNS))
		printKV("TCP", fmt.Sprintf("%d ms", timing.TCP))
		printKV("TLS", fmt.Sprintf("%d ms", timing.TLS))
		printKV("TTFB", fmt.Sprintf("%d ms", timing.TTFB))
		fmt.Println()
	}

	// ▌ Footer
	fmt.Println(Muted(fmt.Sprintf("▌ Completed in %d ms", params.Duration)))
}

func printKV(key, value string) {
	fmt.Printf("  %-10s %s\n", Muted(key), value)
}
