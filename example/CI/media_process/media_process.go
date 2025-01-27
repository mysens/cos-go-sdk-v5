package main

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
)

func log_status(err error) {
	if err == nil {
		return
	}
	if cos.IsNotFoundError(err) {
		// WARN
		fmt.Println("WARN: Resource is not existed")
	} else if e, ok := cos.IsCOSError(err); ok {
		fmt.Printf("ERROR: Code: %v\n", e.Code)
		fmt.Printf("ERROR: Message: %v\n", e.Message)
		fmt.Printf("ERROR: Resource: %v\n", e.Resource)
		fmt.Printf("ERROR: RequestId: %v\n", e.RequestID)
		// ERROR
	} else {
		fmt.Printf("ERROR: %v\n", err)
		// ERROR
	}
}

// InvokeAnimationJob TODO
func InvokeAnimationJob() {
	u, _ := url.Parse("https://test-1234567890.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-1234567890.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "Animation",
		Input: &cos.JobInput{
			Object: "input/test.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			Output: &cos.JobOutput{
				Region: "ap-chongqing",
				Object: "output/game.jpg",
				Bucket: "test-1234567890",
			},
			Animation: &cos.Animation{
				Container: &cos.Container{
					Format: "gif",
				},
				Video: &cos.AnimationVideo{
					Codec:                   "gif",
					AnimateOnlyKeepKeyFrame: "true",
				},
				TimeInterval: &cos.TimeInterval{
					Start:    "0",
					Duration: "",
				},
			},
		},
		QueueId: DescribeQueueRes.QueueList[0].QueueId,
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// InvokeSmartCoverJob TODO
func InvokeSmartCoverJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "SmartCover",
		Input: &cos.JobInput{
			Object: "input/dog.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			SmartCover: &cos.NodeSmartCover{
				Format:           "jpg",
				Height:           "1280",
				Width:            "720",
				Count:            "2",
				DeleteDuplicates: "true",
			},
			Output: &cos.JobOutput{
				Region: "ap-chongqing",
				Object: "output/mc-${number}.jpg",
				Bucket: "test-123456789",
			},
		},
		QueueId: DescribeQueueRes.QueueList[0].QueueId,
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// InvokeSnapshotJob TODO
func InvokeSnapshotJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-beijing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-beijing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "Snapshot",
		Input: &cos.JobInput{
			Object: "input/117374C.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			Output: &cos.JobOutput{
				Region: "ap-beijing",
				Object: "output/abc-${Number}.jpg",
				Bucket: "test-123456789",
			},
			Snapshot: &cos.Snapshot{
				Mode:  "Interval",
				Start: "0",
				Count: "1",
			},
		},
		QueueId: DescribeQueueRes.QueueList[0].QueueId,
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// InvokeConcatJob TODO
func InvokeConcatJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-beijing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-beijing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	concatFragment := make([]cos.ConcatFragment, 0)
	concatFragment = append(concatFragment, cos.ConcatFragment{
		Url:           "https://test-123456789.cos.ap-beijing.myqcloud.com/input/117374C.mp4",
		StartTime:     "0",
		EndTime:       "10",
		FragmentIndex: "0",
	})
	concatFragment = append(concatFragment, cos.ConcatFragment{
		Url:           "https://test-123456789.cos.ap-beijing.myqcloud.com/input/117374C.mp4",
		StartTime:     "20",
		EndTime:       "30",
		FragmentIndex: "1",
	})
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "Concat",
		Operation: &cos.MediaProcessJobOperation{
			Output: &cos.JobOutput{
				Region: "ap-beijing",
				Object: "output/go_117374C.mp4",
				Bucket: "test-123456789",
			},
			ConcatTemplate: &cos.ConcatTemplate{
				Container: &cos.Container{
					Format: "mp4",
				},
				Video: &cos.Video{
					Codec: "H.264",
				},
				Audio: &cos.Audio{
					Codec: "AAC",
				},
				ConcatFragment: concatFragment,
			},
		},
		QueueId: DescribeQueueRes.QueueList[0].QueueId,
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// InvokeTranscodeJob TODO
func InvokeTranscodeJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "Transcode",
		Input: &cos.JobInput{
			Object: "input/117374C.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			Output: &cos.JobOutput{
				Region: "ap-chongqing",
				Object: "output/go_117374C.mp4",
				Bucket: "test-123456789",
			},
			Transcode: &cos.Transcode{
				Container: &cos.Container{
					Format: "mp4",
				},
				Video: &cos.Video{
					Codec: "H.264",
				},
				Audio: &cos.Audio{
					Codec: "AAC",
				},
				TimeInterval: &cos.TimeInterval{
					Start:    "10",
					Duration: "",
				},
			},
			UserData: "hello world",
		},
		QueueId: "paaf4fce5521a40888a3034a5de80f6ca",
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// InvokeMultiJobs TODO
func InvokeMultiJobs() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMultiMediaJobsOptions{
		Input: &cos.JobInput{
			Object: "input/117374C.mp4",
		},
		Operation: []cos.MediaProcessJobOperation{
			cos.MediaProcessJobOperation{
				Tag: "Snapshot",
				Output: &cos.JobOutput{
					Region: "ap-chongqing",
					Object: "output/go_${Number}.mp4",
					Bucket: "test-123456789",
				},
				Snapshot: &cos.Snapshot{
					Mode:  "Interval",
					Start: "0",
					Count: "1",
				},
			},
			cos.MediaProcessJobOperation{
				Tag: "Transcode",
				Output: &cos.JobOutput{
					Region: "ap-chongqing",
					Object: "output/go_117374C.mp4",
					Bucket: "test-123456789",
				},
				Transcode: &cos.Transcode{
					Container: &cos.Container{
						Format: "mp4",
					},
					Video: &cos.Video{
						Codec: "H.264",
					},
					Audio: &cos.Audio{
						Codec: "AAC",
					},
					TimeInterval: &cos.TimeInterval{
						Start:    "10",
						Duration: "",
					},
				},
			},
			cos.MediaProcessJobOperation{
				Tag: "Animation",
				Output: &cos.JobOutput{
					Region: "ap-chongqing",
					Object: "output/go_117374C.gif",
					Bucket: "test-123456789",
				},
				Animation: &cos.Animation{
					Container: &cos.Container{
						Format: "gif",
					},
					Video: &cos.AnimationVideo{
						Codec:                   "gif",
						AnimateOnlyKeepKeyFrame: "true",
					},
					TimeInterval: &cos.TimeInterval{
						Start:    "0",
						Duration: "",
					},
				},
			},
		},
		QueueId: "paaf4fce5521a40888a3034a5de80f6ca",
	}
	createJobRes, _, err := c.CI.CreateMultiMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	for k, job := range createJobRes.JobsDetail {
		fmt.Printf("job:%d, %+v\n", k, job)
	}
}

// JobNotifyCallback TODO
func JobNotifyCallback() {
	taskBody := "<Response><JobsDetail><Code>Success</Code><CreationTime>2022-02-09T11:25:43+0800</CreationTime><EndTime>2022-02-09T11:25:47+0800</EndTime><Input><BucketId>test-123456789</BucketId><Object>input/117374C.mp4</Object><Region>ap-chongqing</Region></Input><JobId>jf6717076895711ecafdd594be6cca70c</JobId><Message/><Operation><MediaInfo><Format><Bitrate>215.817000</Bitrate><Duration>96.931000</Duration><FormatLongName>QuickTime / MOV</FormatLongName><FormatName>mov,mp4,m4a,3gp,3g2,mj2</FormatName><NumProgram>0</NumProgram><NumStream>1</NumStream><Size>2614921</Size><StartTime>0.000000</StartTime></Format><Stream><Audio/><Subtitle/><Video><AvgFps>29.970030</AvgFps><Bitrate>212.875000</Bitrate><CodecLongName>H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10</CodecLongName><CodecName>h264</CodecName><CodecTag>0x31637661</CodecTag><CodecTagString>avc1</CodecTagString><CodecTimeBase>1/30000</CodecTimeBase><Dar>16:9</Dar><Duration>96.930167</Duration><Fps>29.970030</Fps><HasBFrame>2</HasBFrame><Height>400</Height><Index>0</Index><Language>und</Language><Level>30</Level><NumFrames>2905</NumFrames><PixFormat>yuv420p</PixFormat><Profile>High</Profile><RefFrames>1</RefFrames><Rotation>0.000000</Rotation><Sar>640:639</Sar><StartTime>0.000000</StartTime><Timebase>1/30000</Timebase><Width>710</Width></Video></Stream></MediaInfo><MediaResult><OutputFile><Bucket>test-123456789</Bucket><Md5Info><Md5>38f0b40c78562f819421137541043f09</Md5><ObjectName>output/go_117374C.mp4</ObjectName></Md5Info><ObjectName>output/go_117374C.mp4</ObjectName><ObjectPrefix/><Region>ap-chongqing</Region><SpriteOutputFile><Bucket/><Md5Info/><ObjectName/><ObjectPrefix/><Region/></SpriteOutputFile></OutputFile></MediaResult><Output><Bucket>test-123456789</Bucket><Object>output/go_117374C.mp4</Object><Region>ap-chongqing</Region></Output><Transcode><Audio><Bitrate/><Channels/><Codec>AAC</Codec><KeepTwoTracks>false</KeepTwoTracks><Profile/><Remove>false</Remove><SampleFormat/><Samplerate>44100</Samplerate><SwitchTrack>false</SwitchTrack></Audio><Container><Format>mp4</Format></Container><TimeInterval><Duration/><Start>10</Start></TimeInterval><TransConfig><AdjDarMethod/><AudioBitrateAdjMethod/><DeleteMetadata>false</DeleteMetadata><IsCheckAudioBitrate>false</IsCheckAudioBitrate><IsCheckReso>false</IsCheckReso><IsCheckVideoBitrate>false</IsCheckVideoBitrate><IsHdr2Sdr>false</IsHdr2Sdr><IsStreamCopy>false</IsStreamCopy><ResoAdjMethod/><VideoBitrateAdjMethod/></TransConfig><Video><AnimateFramesPerSecond/><AnimateOnlyKeepKeyFrame>false</AnimateOnlyKeepKeyFrame><AnimateTimeIntervalOfFrame/><Bitrate/><Bufsize/><Codec>H.264</Codec><Crf>25</Crf><Crop/><Fps/><Gop/><Height/><Interlaced>false</Interlaced><LongShortMode>false</LongShortMode><Maxrate/><Pad/><Pixfmt/><Preset>medium</Preset><Profile>high</Profile><Quality/><Remove>false</Remove><ScanMode/><SliceTime>5</SliceTime><Width/></Video></Transcode></Operation><Progress>100</Progress><QueueId>paaf4fce5521a40888a3034a5de80f6ca</QueueId><StartTime>2022-02-09T11:25:43+0800</StartTime><State>Success</State><Tag>Transcode</Tag></JobsDetail></Response>"
	var body cos.MediaProcessJobsNotifyBody
	err := xml.Unmarshal([]byte(taskBody), &body)
	if err != nil {
		fmt.Println(fmt.Sprintf("err:%v", err))
	} else {
		fmt.Println(fmt.Sprintf("body:%+v", body))
		fmt.Println(fmt.Sprintf("mediaInfo:%+v", body.JobsDetail.Operation.MediaInfo))
		fmt.Println(fmt.Sprintf("mediaResult:%+v", body.JobsDetail.Operation.MediaResult))
	}
}

// WorkflowExecutionNotifyCallback TODO
func WorkflowExecutionNotifyCallback() {
	workflowExecutionBody := "<Response><EventName>WorkflowFinish</EventName><WorkflowExecution><RunId>i70ae991a152911ecb184525400a8700f</RunId><BucketId></BucketId><Object>62ddbc1245.mp4</Object><CosHeaders><Key>x-cos-meta-id</Key><Value>62ddbc1245</Value></CosHeaders><CosHeaders><Key>Content-Type</Key><Value>video/mp4</Value></CosHeaders><WorkflowId>w29ba54d02b7340dd9fb44eb5beb786b9</WorkflowId><WorkflowName></WorkflowName><CreateTime>2021-09-14 15:00:26+0800</CreateTime><State>Success</State><Tasks><Type>Transcode</Type><CreateTime>2021-09-14 15:00:27+0800</CreateTime><EndTime>2021-09-14 15:00:42+0800</EndTime><State>Success</State><JobId>j70bab192152911ecab79bba409874f7f</JobId><Name>Transcode_1607323983818</Name><TemplateId>t088613dea8d564a9ba7e6b02cbd5de877</TemplateId><TemplateName>HLS-FHD</TemplateName></Tasks></WorkflowExecution></Response>"
	var body cos.WorkflowExecutionNotifyBody
	err := xml.Unmarshal([]byte(workflowExecutionBody), &body)
	if err != nil {
		fmt.Println(fmt.Sprintf("err:%v", err))
	} else {
		fmt.Println(fmt.Sprintf("body:%v", body))
	}
}

// InvokeSpriteSnapshotJob TODO
func InvokeSpriteSnapshotJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "Snapshot",
		Input: &cos.JobInput{
			Object: "input/117374C.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			Output: &cos.JobOutput{
				Region:       "ap-chongqing",
				Object:       "output/abc-${Number}.jpg",
				Bucket:       "test-123456789",
				SpriteObject: "output/sprite-${Number}.jpg",
			},
			Snapshot: &cos.Snapshot{
				Mode:            "Interval",
				Start:           "0",
				Count:           "100",
				SnapshotOutMode: "SnapshotAndSprite", // OnlySnapshot OnlySprite
				SpriteSnapshotConfig: &cos.SpriteSnapshotConfig{
					CellHeight: "128",
					CellWidth:  "128",
					Color:      "Black",
					Columns:    "3",
					Lines:      "10",
					Margin:     "2",
				},
			},
		},
		QueueId: DescribeQueueRes.QueueList[0].QueueId,
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// InvokeSegmentJob TODO
func InvokeSegmentJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "Segment",
		Input: &cos.JobInput{
			Object: "input/117374C.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			Output: &cos.JobOutput{
				Region: "ap-chongqing",
				Object: "output/m3u8/a",
				Bucket: "test-123456789",
			},
			Segment: &cos.Segment{
				Format:   "hls",
				Duration: "10",
				HlsEncrypt: &cos.HlsEncrypt{
					IsHlsEncrypt: true,
					UriKey:       "http://abc.com/",
				},
			},
		},
		QueueId: DescribeQueueRes.QueueList[0].QueueId,
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// DescribeMultiMediaJob TODO
func DescribeMultiMediaJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "Segment",
		Input: &cos.JobInput{
			Object: "input/117374C.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			Output: &cos.JobOutput{
				Region: "ap-chongqing",
				Object: "output/abc-${Number}.mp4",
				Bucket: "test-123456789",
			},
			Segment: &cos.Segment{
				Format:   "mp4",
				Duration: "10",
			},
		},
		QueueId: DescribeQueueRes.QueueList[0].QueueId,
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	jobids := make([]string, 0)
	jobids = append(jobids, createJobRes.JobsDetail.JobId)
	jobids = append(jobids, "a")
	jobids = append(jobids, "b")
	jobids = append(jobids, "c")
	DescribeJobRes, _, err := c.CI.DescribeMultiMediaJob(context.Background(), jobids)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// GetPrivateM3U8 TODO
func GetPrivateM3U8() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	getPrivateM3U8Opt := &cos.GetPrivateM3U8Options{
		Expires: 3600,
	}
	getPrivateM3U8Res, err := c.CI.GetPrivateM3U8(context.Background(), "output/linkv.m3u8", getPrivateM3U8Opt)
	log_status(err)
	fmt.Printf("%+v\n", getPrivateM3U8Res)
}

// InvokeVideoMontageJob TODO
func InvokeVideoMontageJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "VideoMontage",
		Input: &cos.JobInput{
			Object: "input/117374C.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			Output: &cos.JobOutput{
				Region: "ap-chongqing",
				Object: "output/go_117374C.mp4",
				Bucket: "test-123456789",
			},
			VideoMontage: &cos.VideoMontage{
				Container: &cos.Container{
					Format: "mp4",
				},
				Video: &cos.VideoMontageVideo{
					Codec: "H.264",
				},
				Audio: &cos.Audio{
					Codec: "AAC",
				},
			},
		},
		QueueId: "paaf4fce5521a40888a3034a5de80f6ca",
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// InvokeVoiceSeparateJob TODO
func InvokeVoiceSeparateJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "VoiceSeparate",
		Input: &cos.JobInput{
			Object: "example.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			Output: &cos.JobOutput{
				Region: "ap-chongqing",
				Object: "output/go_example.mp4",
				Bucket: "test-123456789",
			},
			VoiceSeparate: &cos.VoiceSeparate{
				AudioMode: "AudioAndBackground",
				AudioConfig: &cos.AudioConfig{
					Codec: "AAC",
				},
			},
		},
		QueueId: "paaf4fce5521a40888a3034a5de80f6ca",
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// InvokeVideoProcessJob TODO
func InvokeVideoProcessJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "VideoProcess",
		Input: &cos.JobInput{
			Object: "example.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			Output: &cos.JobOutput{
				Region: "ap-chongqing",
				Object: "output/vp_example.mp4",
				Bucket: "test-123456789",
			},
			Transcode: &cos.Transcode{
				Container: &cos.Container{
					Format: "mp4",
				},
				Video: &cos.Video{
					Codec: "H.264",
				},
				Audio: &cos.Audio{
					Codec: "AAC",
				},
			},
			VideoProcess: &cos.VideoProcess{
				ColorEnhance: &cos.ColorEnhance{
					Enable: "true",
				},
				MsSharpen: &cos.MsSharpen{
					Enable: "true",
				},
			},
		},
		QueueId: "paaf4fce5521a40888a3034a5de80f6ca",
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// InvokeSDRtoHDRJob TODO
func InvokeSDRtoHDRJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "SDRtoHDR",
		Input: &cos.JobInput{
			Object: "linkv.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			Output: &cos.JobOutput{
				Region: "ap-chongqing",
				Object: "output/sdrtohdr_linkv.mp4",
				Bucket: "test-123456789",
			},
			Transcode: &cos.Transcode{
				Container: &cos.Container{
					Format: "mp4",
				},
				Video: &cos.Video{
					Codec: "H.265",
				},
				Audio: &cos.Audio{
					Codec: "AAC",
				},
			},
			SDRtoHDR: &cos.SDRtoHDR{
				HdrMode: "HLG",
			},
		},
		QueueId: "paaf4fce5521a40888a3034a5de80f6ca",
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// InvokeSuperResolutionJob TODO
func InvokeSuperResolutionJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "SuperResolution",
		Input: &cos.JobInput{
			Object: "100986-2999.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			Output: &cos.JobOutput{
				Region: "ap-chongqing",
				Object: "output/sp-100986-2999.mp4",
				Bucket: "test-123456789",
			},
			Transcode: &cos.Transcode{
				Container: &cos.Container{
					Format: "mp4",
				},
				Video: &cos.Video{
					Codec: "H.264",
				},
				Audio: &cos.Audio{
					Codec: "AAC",
				},
			},
			SuperResolution: &cos.SuperResolution{
				Resolution:    "hdto4k",
				EnableScaleUp: "true",
			},
		},
		QueueId: "paaf4fce5521a40888a3034a5de80f6ca",
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// TriggerWorkflow TODO
func TriggerWorkflow() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	triggerWorkflowOpt := &cos.TriggerWorkflowOptions{
		WorkflowId: "w18fd791485904afba3ab07ed57d9cf1e",
		Object:     "100986-2999.mp4",
	}
	triggerWorkflowRes, _, err := c.CI.TriggerWorkflow(context.Background(), triggerWorkflowOpt)
	log_status(err)
	fmt.Printf("%+v\n", triggerWorkflowRes)
}

// DescribeWorkflowExecutions TODO
func DescribeWorkflowExecutions() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	describeWorkflowExecutionsOpt := &cos.DescribeWorkflowExecutionsOptions{
		WorkflowId: "w18fd791485904afba3ab07ed57d9cf1e",
	}
	describeWorkflowExecutionsRes, _, err := c.CI.DescribeWorkflowExecutions(context.Background(), describeWorkflowExecutionsOpt)
	log_status(err)
	fmt.Printf("%+v\n", describeWorkflowExecutionsRes)
}

// DescribeMultiWorkflowExecution TODO
func DescribeMultiWorkflowExecution() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	describeWorkflowExecutionsRes, _, err := c.CI.DescribeWorkflowExecution(context.Background(), "i00689df860ad11ec9c5952540019ee59")
	log_status(err)
	a, _ := json.Marshal(describeWorkflowExecutionsRes)
	fmt.Println(string(a))
	fmt.Printf("%+v\n", describeWorkflowExecutionsRes)
}

// InvokeASRJob TODO
func InvokeASRJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// CreateMediaJobs
	createJobOpt := &cos.CreateASRJobsOptions{
		Tag: "SpeechRecognition",
		Input: &cos.JobInput{
			Object: "abc.mp3",
		},
		Operation: &cos.ASRJobOperation{
			Output: &cos.JobOutput{
				Region: "ap-chongqing",
				Object: "music.txt",
				Bucket: "test-123456789",
			},
			SpeechRecognition: &cos.SpeechRecognition{
				ChannelNum:      "1",
				EngineModelType: "8k_zh",
			},
		},
		QueueId: "p1db6a1a76ff04806b6af0d96e9bc80ab",
	}
	createJobRes, _, err := c.CI.CreateASRJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)
	DescribeJobRes, _, err := c.CI.DescribeMultiASRJob(context.Background(), []string{createJobRes.JobsDetail.JobId})
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// DescribeASRJob TODO
func DescribeASRJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	DescribeJobRes, _, err := c.CI.DescribeMultiASRJob(context.Background(), []string{"sa59de0a06a4711ec9b81df13272c69a9"})
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail[0].Operation.SpeechRecognitionResult)
}

// DescribeJob TODO
func DescribeJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), "a1e4825762f3f11edafa8d13569009b8c")
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// DescribeJobs TODO
func DescribeJobs() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	opt := &cos.DescribeMediaJobsOptions{
		QueueId: "pa27b2bd96bef43b6baba820175485532",
		Tag:     "Transcode",
	}
	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJobs(context.Background(), opt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// GenerateMediaInfo TODO
func GenerateMediaInfo() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	opt := &cos.GenerateMediaInfoOptions{
		Input: &cos.JobInput{
			Object: "input/117374C.mp4",
		},
	}
	// DescribeMediaJobs
	res, _, err := c.CI.GenerateMediaInfo(context.Background(), opt)
	log_status(err)
	fmt.Printf("%+v\n", res)
}

// InvokeMediaInfoJob TODO
func InvokeMediaInfoJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "MediaInfo",
		Input: &cos.JobInput{
			Object: "input/117374C.mp4",
		},
		QueueId: DescribeQueueRes.QueueList[0].QueueId,
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	for {
		time.Sleep(100 * time.Second)
		// DescribeMediaJobs
		DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
		log_status(err)
		fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
		if DescribeJobRes.JobsDetail.State == "Success" {
			break
		}
	}
}

// InvokeStreamExtractJob TODO
func InvokeStreamExtractJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	streamEtract := make([]cos.StreamExtract, 0)
	streamEtract = append(streamEtract, cos.StreamExtract{
		Index:  "1",
		Object: "stream/video02_1.mp4",
	})

	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "StreamExtract",
		Input: &cos.JobInput{
			Object: "video02.mp4",
		},
		QueueId: DescribeQueueRes.QueueList[0].QueueId,
		Operation: &cos.MediaProcessJobOperation{
			Output: &cos.JobOutput{
				Region:        "ap-chongqing",
				Bucket:        "test-123456789",
				StreamExtract: streamEtract,
			},
		},
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	for {
		time.Sleep(100 * time.Second)
		// DescribeMediaJobs
		DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
		log_status(err)
		fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
		if DescribeJobRes.JobsDetail.State == "Success" {
			break
		}
	}
}

// InvokePicProcessJob TODO
func InvokePicProcessJob() {
	// todo 需要替换为自己的域名
	u, _ := url.Parse("https://testpic-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://testpic-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// todo 需要替换为自己的secretid  secretkey
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribePicProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribePicProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreatePicJobsOptions{
		Tag: "PicProcess",
		Input: &cos.JobInput{
			// todo 需要替换为自己的Input文件
			Object: "1.png",
		},
		QueueId: DescribeQueueRes.QueueList[0].QueueId,
		Operation: &cos.PicProcessJobOperation{
			// todo 需要替换为自己的图片处理配置
			PicProcess: &cos.PicProcess{
				IsPicInfo:   "true",
				ProcessRule: "imageMogr2/format/jpg/interlace/0/quality/100",
			},
			// todo 需要替换为自己的Output信息
			Output: &cos.JobOutput{
				Region: "ap-chongqing",
				Bucket: "testpic-123456789",
				Object: "test.jpg",
			},
		},
		// todo 需要替换为自己的回调地址信息
		CallBack: "https://demo.org/callback",
	}
	createJobRes, _, err := c.CI.CreatePicProcessJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	for {
		time.Sleep(2 * time.Second)
		// DescribeMediaJobs
		DescribeJobRes, _, err := c.CI.DescribePicProcessJob(context.Background(), createJobRes.JobsDetail.JobId)
		log_status(err)
		fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
		if DescribeJobRes.JobsDetail.State == "Success" {
			break
		}
	}
}

// InvokeDigitalWatermarkJob TODO
func InvokeDigitalWatermarkJob() {
	u, _ := url.Parse("https://test-1234567890.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-1234567890.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "DigitalWatermark",
		Input: &cos.JobInput{
			Object: "input/test.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			Output: &cos.JobOutput{
				Region: "ap-chongqing",
				Object: "output/test.mp4",
				Bucket: "test-1234567890",
			},
			DigitalWatermark: &cos.DigitalWatermark{
				Message: "HelloWorld",
				Type:    "Text",
				Version: "V1",
			},
		},
		QueueId: DescribeQueueRes.QueueList[0].QueueId,
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// InvokeExtractDigitalWatermarkJob TODO
func InvokeExtractDigitalWatermarkJob() {
	u, _ := url.Parse("https://test-1234567890.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-1234567890.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "ExtractDigitalWatermark",
		Input: &cos.JobInput{
			Object: "output/test.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			ExtractDigitalWatermark: &cos.ExtractDigitalWatermark{
				Type:    "Text",
				Version: "V1",
			},
		},
		QueueId: DescribeQueueRes.QueueList[0].QueueId,
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// InvokeVideoTagJob TODO
func InvokeVideoTagJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "VideoTag",
		Input: &cos.JobInput{
			Object: "input/dog.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			VideoTag: &cos.VideoTag{
				Scenario: "Stream",
			},
		},
		QueueId: DescribeQueueRes.QueueList[0].QueueId,
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// InvokeNoiseReductionJob TODO
func InvokeNoiseReductionJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "NoiseReduction",
		Input: &cos.JobInput{
			Object: "input/zhanghuimei_wen.mp3",
		},
		Operation: &cos.MediaProcessJobOperation{
			Output: &cos.JobOutput{
				Region: "ap-chongqing",
				Bucket: "test-123456789",
				Object: "output/out.mp3",
			},
			UserData: "This is my NoiseReduction job",
			JobLevel: 1,
		},
		QueueId: DescribeQueueRes.QueueList[0].QueueId,
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// InvokeQualityEstimateJob TODO
func InvokeQualityEstimateJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "QualityEstimate",
		Input: &cos.JobInput{
			Object: "input/dog.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			UserData: "This is my QualityEstimate job",
			JobLevel: 1,
			QualityEstimateConfig: &cos.QualityEstimateConfig{
				Rotate: "90", // 只支持0 90 180 270
			},
		},
		QueueId: DescribeQueueRes.QueueList[0].QueueId,
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// InvokeTtsJob TODO
func InvokeTtsJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	DescribeQueueRes, _, err := c.CI.DescribeMediaProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "Tts",
		Operation: &cos.MediaProcessJobOperation{
			TtsTpl: &cos.TtsTpl{
				Mode:      "Sync",
				Codec:     "mp3",
				VoiceType: "aixiaonan",
				Volume:    "5",
				Speed:     "150",
			},
			TtsConfig: &cos.TtsConfig{
				Input:     "床前明月光，疑是地上霜",
				InputType: "Text",
			},
			Output: &cos.JobOutput{
				Region: "ap-chongqing",
				Bucket: "test-123456789",
				Object: "output/out.mp3",
			},
			UserData: "This is my Tts job",
			JobLevel: 1,
		},
		QueueId: DescribeQueueRes.QueueList[0].QueueId,
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// InvokeTranslationJob TODO
func InvokeTranslationJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	// 注意这里是AI
	DescribeQueueRes, _, err := c.CI.DescribeAIProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateAIJobsOptions{
		Tag: "Translation",
		Input: &cos.JobInput{
			Object:    "input/translate-en.txt",
			Lang:      "en",
			Type:      "txt",
			BasicType: "",
		},
		Operation: &cos.MediaProcessJobOperation{
			Translation: &cos.Translation{
				Lang: "zh",
				Type: "txt",
			},
			Output: &cos.JobOutput{
				Region: "ap-chongqing",
				Bucket: "test-123456789",
				Object: "output/out.txt",
			},
			UserData: "This is my Translation job",
			JobLevel: 1,
		},
		QueueId: DescribeQueueRes.QueueList[0].QueueId,
	}
	// 注意这里是AI
	createJobRes, _, err := c.CI.CreateAIJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeAIJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// InvokeWordsGeneralizeJob TODO
func InvokeWordsGeneralizeJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// DescribeMediaProcessQueues
	DescribeQueueOpt := &cos.DescribeMediaProcessQueuesOptions{
		QueueIds:   "",
		PageNumber: 1,
		PageSize:   2,
	}
	// 注意这里是AI
	DescribeQueueRes, _, err := c.CI.DescribeAIProcessQueues(context.Background(), DescribeQueueOpt)
	log_status(err)
	fmt.Printf("%+v\n", DescribeQueueRes)
	// CreateMediaJobs
	createJobOpt := &cos.CreateAIJobsOptions{
		Tag: "WordsGeneralize",
		Input: &cos.JobInput{
			Object: "input/WordsGeneralize.txt",
		},
		Operation: &cos.MediaProcessJobOperation{
			WordsGeneralize: &cos.WordsGeneralize{
				NerMethod: "DL",
				SegMethod: "MIX",
			},
			Output: &cos.JobOutput{
				Region: "ap-chongqing",
				Bucket: "test-123456789",
				Object: "output/out.txt",
			},
			UserData: "This is my WordsGeneralize job",
			JobLevel: 1,
		},
		QueueId: DescribeQueueRes.QueueList[0].QueueId,
	}
	// 注意这里是AI
	createJobRes, _, err := c.CI.CreateAIJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeAIJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// InvokeMediaInfoJob TODO
func PostSnapshot() {
	u, _ := url.Parse("https://test-1234567890.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-1234567890.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// PostSnapshotOptions
	PostSnapshotOpt := &cos.PostSnapshotOptions{
		Input: &cos.JobInput{
			Object: "example.mp4",
		},
		Time:   "1",
		Width:  128,
		Height: 128,
		Format: "png",
		Output: &cos.JobOutput{
			Region: "ap-chongqing",
			Bucket: "test-1234567890",
			Object: "example.mp4.png",
		},
	}
	PostSnapshotRes, _, err := c.CI.PostSnapshot(context.Background(), PostSnapshotOpt)
	log_status(err)
	fmt.Printf("%+v\n", PostSnapshotRes)
}

func InvokeSplitVideoPartsJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// CreateMediaJobs
	createJobOpt := &cos.CreateAIJobsOptions{
		Tag: "SplitVideoParts",
		Input: &cos.JobInput{
			Object: "input/test.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			SplitVideoParts: &cos.SplitVideoParts{
				Mode: "SHOTDETECT",
			},
			UserData: "This is my SplitVideoParts job",
			JobLevel: 1,
		},
	}
	// 注意这里是AI
	createJobRes, _, err := c.CI.CreateAIJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeAIJob
	DescribeJobRes, _, err := c.CI.DescribeAIJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

func InvokeVideoEnhanceJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// CreateMediaJobs
	createJobOpt := &cos.CreateMediaJobsOptions{
		Tag: "VideoEnhance",
		Input: &cos.JobInput{
			Object: "input/test.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			VideoEnhance: &cos.VideoEnhance{
				Transcode: &cos.Transcode{
					Container: &cos.Container{
						Format: "mp4",
					},
					Video: &cos.Video{
						Codec:   "H.264",
						Bitrate: "1000",
						Width:   "1280",
						Fps:     "30",
					},
					Audio: &cos.Audio{
						Codec:      "aac",
						Bitrate:    "128",
						Samplerate: "44100",
						Channels:   "4",
					},
				},
				SuperResolution: &cos.SuperResolution{
					Resolution:    "sdtohd",
					EnableScaleUp: "true",
					Version:       "Enhance",
				},
				ColorEnhance: &cos.ColorEnhance{
					Contrast:   "50",
					Correction: "100",
					Saturation: "100",
				},
				MsSharpen: &cos.MsSharpen{
					SharpenLevel: "5",
				},
				SDRtoHDR: &cos.SDRtoHDR{
					HdrMode: "HDR10",
				},
				FrameEnhance: &cos.FrameEnhance{
					FrameDoubling: "true",
				},
			},
			UserData: "This is my SplitVideoParts job",
			JobLevel: 1,
		},
	}
	createJobRes, _, err := c.CI.CreateMediaJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeAIJob
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

func InvokeVideoTargetRecJob() {
	u, _ := url.Parse("https://test-1234567890.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-1234567890.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// CreateMediaJobs
	createJobOpt := &cos.CreateAIJobsOptions{
		Tag: "VideoTargetRec",
		Input: &cos.JobInput{
			Object: "input/test.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			VideoTargetRec: &cos.VideoTargetRec{
				Body: "true",
				Pet:  "true",
				Car:  "true",
			},
			UserData: "This is my VideoTargetRec job",
			JobLevel: 1,
		},
	}
	// 注意这里是AI
	createJobRes, _, err := c.CI.CreateAIJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeAIJob
	DescribeJobRes, _, err := c.CI.DescribeAIJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

func InvokeSegmentVideoBodyJob() {
	u, _ := url.Parse("https://test-1234567890.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-1234567890.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// CreateMediaJobs
	createJobOpt := &cos.CreateAIJobsOptions{
		Tag: "SegmentVideoBody",
		Input: &cos.JobInput{
			Object: "input/test.mp4",
		},
		Operation: &cos.MediaProcessJobOperation{
			SegmentVideoBody: &cos.SegmentVideoBody{
				Mode: "Mask",
			},
			UserData: "This is my SegmentVideoBody job",
			JobLevel: 1,
		},
	}
	// 注意这里是AI
	createJobRes, _, err := c.CI.CreateAIJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeAIJob
	DescribeJobRes, _, err := c.CI.DescribeAIJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

// InvokeSoundHoundJob TODO
func InvokeSoundHoundJob() {
	u, _ := url.Parse("https://test-123456789.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-123456789.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// CreateMediaJobs
	createJobOpt := &cos.CreateASRJobsOptions{
		Tag: "SoundHound",
		Operation: &cos.ASRJobOperation{
			UserData: "This is my SoundHound job",
			JobLevel: 1,
		},
	}
	createJobRes, _, err := c.CI.CreateASRJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)

	// DescribeMediaJobs
	DescribeJobRes, _, err := c.CI.DescribeMediaJob(context.Background(), createJobRes.JobsDetail.JobId)
	log_status(err)
	fmt.Printf("%+v\n", DescribeJobRes.JobsDetail)
}

func InvokePosterProductionJob() {
	u, _ := url.Parse("https://test-1234567890.cos.ap-chongqing.myqcloud.com")
	cu, _ := url.Parse("https://test-1234567890.ci.ap-chongqing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// todo 需要替换为自己的secretid  secretkey
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	// CreatePicProcessJobs
	type autoInfo struct {
		TextMain string `xml:"text_main,omitempty"`
		TextSub  string `xml:"text_sub,omitempty"`
	}
	info := autoInfo{
		TextMain: "父亲节快乐",
		TextSub:  "献给最伟大的父亲!!!",
	}
	createJobOpt := &cos.CreatePicJobsOptions{
		Tag: "PosterProduction",
		Operation: &cos.PicProcessJobOperation{
			PosterProduction: &cos.PosterProduction{
				TemplateId: "6444f12ae24d596cdbd774fb",
				Info:       info,
			},
			Output: &cos.JobOutput{
				Region: "ap-chongqing",
				Bucket: "test-1234567890",
				Object: "poster/PosterProduction2.jpg",
			},
		},
		// todo 需要替换为自己的回调地址信息
		CallBack: "https://demo.org/callback",
	}
	createJobRes, _, err := c.CI.CreatePicProcessJobs(context.Background(), createJobOpt)
	log_status(err)
	fmt.Printf("%+v\n", createJobRes.JobsDetail)
}

func main() {
	// InvokeAnimationJob()
	// InvokeSnapshotJob()
	// InvokeSmartCoverJob()
	// InvokeConcatJob()
	// InvokeTranscodeJob()
	// InvokeMultiJobs()
	// JobNotifyCallback()
	// WorkflowExecutionNotifyCallback()
	// InvokeSpriteSnapshotJob()
	// InvokeSegmentJob()
	// DescribeMultiMediaJob()
	// GetPrivateM3U8()
	// InvokeVideoMontageJob()
	// InvokeVoiceSeparateJob()
	// InvokeVideoProcessJob()
	// InvokeSDRtoHDRJob()
	// InvokeSuperResolutionJob()
	// TriggerWorkflow()
	// DescribeWorkflowExecutions()
	// DescribeMultiWorkflowExecution()
	// InvokeASRJob()
	// DescribeASRJob()
	// DescribeJob()
	// DescribeJobs()
	// GenerateMediaInfo()
	// InvokeMediaInfoJob()
	// InvokeStreamExtractJob()
	// InvokePicProcessJob()
	// InvokeDigitalWatermarkJob()
	// InvokeExtractDigitalWatermarkJob()
	// InvokeVideoTagJob()

	// InvokeNoiseReductionJob()
	// InvokeQualityEstimateJob()
	// InvokeTtsJob()
	// InvokeTranslationJob()
	// InvokeWordsGeneralizeJob()
	// PostSnapshot()
	// InvokeVideoEnhanceJob()
	// InvokeSplitVideoPartsJob()
	// InvokeVideoTargetRecJob()
	// InvokeSegmentVideoBodyJob()
	InvokePosterProductionJob()
}
