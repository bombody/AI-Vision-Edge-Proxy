export interface StreamProcess {
    name?:string
    image_tag?:string
    rtsp_endpoint?:string
    rtmp_endpoint?:string
    container_id?:string
    status?:string
    state?:State
    logs?:Logs
    created?:Number
    modified?:Number
    rtmp_stream_status?:RTMPStreamStatus
    upgrade_available?:boolean
    newer_version?:string
    upgrad