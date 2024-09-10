import {defineCodeRunnersSetup} from '@slidev/types'

export default defineCodeRunnersSetup(() => {
    return {
        async go(code, ctx) {
            const result = await executeGoCodeRemotely(code)
            if (result.compilation_error) {
                return {
                    error: "Compilation error: " + result.compilation_error
                }
            }
            if (result.stderr) {
                return {
                    error: "Stderr: " + result.stderr
                }
            }
            return {
                html: result.stdout.replaceAll("\n", "<br>"),
            }
        }
    }
})

interface GoResponse {
    stdout: string
    stderr: string
    compilation_error: string
}

async function executeGoCodeRemotely(code: string): Promise<GoResponse> {
    let request = new Request("http://localhost:8080/", {
        method: "POST",
        body: code,
        cache: "no-cache",
    });
    return fetch(request)
        .then(response => {
            return response.json()
        });
}
