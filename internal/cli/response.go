package cli

import (
	"github.com/sahshad/apix/internal/types"
	"fmt"
	"strings"
)

func RenderResponse(params types.ResponseParams, verbose bool) {

	fmt.Println(Cyan("▌ REQUEST"))
	printKV("Method", params.Method)
	printKV("Endpoint", params.Endpoint)

	fmt.Println()

	fmt.Println(Cyan("▌ RESPONSE"))
	printKV("Status", FormatStatus(params.Status))
	printKV("Duration", fmt.Sprintf("%d ms", params.Duration))
	printKV("Size", params.Size)

	if params.ContentType != "" {
		printKV("Type", params.ContentType)
	}

	fmt.Println()

	if verbose && len(params.Headers) > 0 {
		fmt.Println(Cyan("▌ HEADERS"))
		for k, v := range params.Headers {
			fmt.Printf("  %s: %s\n", strings.ToLower(k), strings.Join(v, ", "))
		}
		fmt.Println()
	}

	fmt.Println(Cyan("▌ BODY"))
	fmt.Println(formatBody(params.Body, params.ContentType))
	fmt.Println()

	timing := params.Timing

	if verbose && timing != nil {
		fmt.Println(Cyan("▌ TIMING"))
		printKV("DNS", fmt.Sprintf("%d ms", timing.DNS))
		printKV("TCP", fmt.Sprintf("%d ms", timing.TCP))
		printKV("TLS", fmt.Sprintf("%d ms", timing.TLS))
		printKV("TTFB", fmt.Sprintf("%d ms", timing.TTFB))
		fmt.Println()
	}

	fmt.Println(Muted(fmt.Sprintf("▌ Completed in %d ms", params.Duration)))
}

func printKV(key, value string) {
	fmt.Printf("  %-10s %s\n", Muted(key), value)
}
