# Fast-Chat

## Design

 * UI provided by Web browser.
   - https://subvisual.co/blog/posts/39-tutorial-html-audio-capture-streaming-to-node-js-no-browser-
   - https://developer.mozilla.org/en-US/docs/Web/API/ScriptProcessorNode
 * Communicate from local web browser to local Go app.
   - Connection via web sockets, using binary proto for voice / video / files, JSON for text.
 * Go client talks to Go server
   - Communicate via gRPC
   - Hash files to be uploaded with blake2b https://github.com/minio/blake2b-simd .
 * Go server maintains centeral database of all activity
   - Use CRDB for backend
   - Messages are persistant and encrypted at rest.
   - Group messages are routed to all recipiants.
   - Messages must be acked, or they will still retain un-read per user

