package cos

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/clbanning/mxj"
	"github.com/mitchellh/mapstructure"
)

// JobInput TODO
type JobInput struct {
	Object string `xml:"Object,omitempty"`
}

// StreamExtract TODO
type StreamExtract struct {
	Index  string `xml:"Index,omitempty"`
	Object string `xml:"Object,omitempty"`
}

// JobOutput TODO
type JobOutput struct {
	Region        string          `xml:"Region,omitempty"`
	Bucket        string          `xml:"Bucket,omitempty"`
	Object        string          `xml:"Object,omitempty"`
	SpriteObject  string          `xml:"SpriteObject,omitempty"`
	AuObject      string          `xml:"AuObject,omitempty"`
	StreamExtract []StreamExtract `xml:"StreamExtract,omitempty"`
}

// Container TODO
type Container struct {
	Format string `xml:"Format"`
}

// Video TODO
type Video struct {
	Codec         string `xml:"Codec"`
	Width         string `xml:"Width,omitempty"`
	Height        string `xml:"Height,omitempty"`
	Fps           string `xml:"Fps,omitempty"`
	Remove        string `xml:"Remove,omitempty"`
	Profile       string `xml:"Profile,omitempty"`
	Bitrate       string `xml:"Bitrate,omitempty"`
	Crf           string `xml:"Crf,omitempty"`
	Gop           string `xml:"Gop,omitempty"`
	Preset        string `xml:"Preset,omitempty"`
	Bufsize       string `xml:"Bufsize,omitempty"`
	Maxrate       string `xml:"Maxrate,omitempty"`
	HlsTsTime     string `xml:"HlsTsTime,omitempty"`
	DashSegment   string `xml:"DashSegment,omitempty"`
	Pixfmt        string `xml:"Pixfmt,omitempty"`
	LongShortMode string `xml:"LongShortMode,omitempty"`
	Rotate        string `xml:"Rotate,omitempty"`
}

// TimeInterval TODO
type TimeInterval struct {
	Start    string `xml:"Start,omitempty"`
	Duration string `xml:"Duration,omitempty"`
}

// Audio TODO
type Audio struct {
	Codec         string `xml:"Codec,omitempty"`
	Samplerate    string `xml:"Samplerate,omitempty"`
	Bitrate       string `xml:"Bitrate,omitempty"`
	Channels      string `xml:"Channels,omitempty"`
	Remove        string `xml:"Remove,omitempty"`
	KeepTwoTracks string `xml:"KeepTwoTracks,omitempty"`
	SwitchTrack   string `xml:"SwitchTrack,omitempty"`
	SampleFormat  string `xml:"SampleFormat,omitempty"`
}

// TransConfig TODO
type TransConfig struct {
	AdjDarMethod          string      `xml:"AdjDarMethod,omitempty"`
	IsCheckReso           string      `xml:"IsCheckReso,omitempty"`
	ResoAdjMethod         string      `xml:"ResoAdjMethod,omitempty"`
	IsCheckVideoBitrate   string      `xml:"IsCheckVideoBitrate,omitempty"`
	VideoBitrateAdjMethod string      `xml:"VideoBitrateAdjMethod,omitempty"`
	IsCheckAudioBitrate   string      `xml:"IsCheckAudioBitrate,omitempty"`
	AudioBitrateAdjMethod string      `xml:"AudioBitrateAdjMethod,omitempty"`
	DeleteMetadata        string      `xml:"DeleteMetadata,omitempty"`
	IsHdr2Sdr             string      `xml:"IsHdr2Sdr,omitempty"`
	HlsEncrypt            *HlsEncrypt `xml:"HlsEncrypt,omitempty"`
}

// Transcode TODO
type Transcode struct {
	Container    *Container    `xml:"Container,omitempty"`
	Video        *Video        `xml:"Video,omitempty"`
	TimeInterval *TimeInterval `xml:"TimeInterval,omitempty"`
	Audio        *Audio        `xml:"Audio,omitempty"`
	TransConfig  *TransConfig  `xml:"TransConfig,omitempty"`
}

// Image TODO
type Image struct {
	Url          string `xml:"Url,omitempty"`
	Mode         string `xml:"Mode,omitempty"`
	Width        string `xml:"Width,omitempty"`
	Height       string `xml:"Height,omitempty"`
	Transparency string `xml:"Transparency,omitempty"`
	Background   string `xml:"Background,omitempty"`
}

// Text TODO
type Text struct {
	FontSize     string `xml:"FontSize,omitempty"`
	FontType     string `xml:"FontType,omitempty"`
	FontColor    string `xml:"FontColor,omitempty"`
	Transparency string `xml:"Transparency,omitempty"`
	Text         string `xml:"Text,omitempty"`
}

// Watermark TODO
type Watermark struct {
	Type      string `xml:"Type,omitempty"`
	Pos       string `xml:"Pos,omitempty"` // TopLeft：左上; Top：上居中; TopRight：右上; Left：左居中; Center：正中心; Right：右居中; BottomLeft：左下; Bottom：下居中; BottomRight：右下
	LocMode   string `xml:"LocMode,omitempty"`
	Dx        string `xml:"Dx,omitempty"`
	Dy        string `xml:"Dy,omitempty"`
	StartTime string `xml:"StartTime,omitempty"`
	EndTime   string `xml:"EndTime,omitempty"`
	Image     *Image `xml:"Image,omitempty"`
	Text      *Text  `xml:"Text,omitempty"`
}

// EffectConfig TODO
type EffectConfig struct {
	EnableStartFadein string `xml:"EnableStartFadein,omitempty"`
	StartFadeinTime   string `xml:"StartFadeinTime,omitempty"`
	EnableEndFadeout  string `xml:"EnableEndFadeout,omitempty"`
	EndFadeoutTime    string `xml:"EndFadeoutTime,omitempty"`
	EnableBgmFade     string `xml:"EnableBgmFade,omitempty"`
	BgmFadeTime       string `xml:"BgmFadeTime,omitempty"`
}

// AudioMix TODO
type AudioMix struct {
	AudioSource  string        `xml:"AudioSource,omitempty"`
	MixMode      string        `xml:"MixMode,omitempty"`
	Replace      string        `xml:"Replace,omitempty"`
	EffectConfig *EffectConfig `xml:"EffectConfig,omitempty"`
}

// ConcatFragment TODO
type ConcatFragment struct {
	Url       string `xml:"Url,omitempty"`
	Mode      string `xml:"Mode,omitempty"`
	StartTime string `xml:"StartTime,omitempty"`
	EndTime   string `xml:"EndTime,omitempty"`
}

// ConcatTemplate TODO
type ConcatTemplate struct {
	ConcatFragment []ConcatFragment `xml:"ConcatFragment,omitempty"`
	Audio          *Audio           `xml:"Audio,omitempty"`
	Video          *Video           `xml:"Video,omitempty"`
	Container      *Container       `xml:"Container,omitempty"`
	Index          string           `xml:"Index,omitempty"`
	AudioMix       *AudioMix        `xml:"AudioMix,omitempty"`
}

// SpriteSnapshotConfig TODO
type SpriteSnapshotConfig struct {
	CellHeight string `xml:"CellHeight,omitempty"`
	CellWidth  string `xml:"CellWidth,omitempty"`
	Color      string `xml:"Color,omitempty"`
	Columns    string `xml:"Columns,omitempty"`
	Lines      string `xml:"Lines,omitempty"`
	Margin     string `xml:"Margin,omitempty"`
	Padding    string `xml:"Padding,omitempty"`
}

// Snapshot TODO
type Snapshot struct {
	Mode                 string                `xml:"Mode,omitempty"`
	Start                string                `xml:"Start,omitempty"`
	TimeInterval         string                `xml:"TimeInterval,omitempty"`
	Count                string                `xml:"Count,omitempty"`
	Width                string                `xml:"Width,omitempty"`
	Height               string                `xml:"Height,omitempty"`
	CIParam              string                `xml:"CIParam,omitempty"`
	IsCheckCount         bool                  `xml:"IsCheckCount,omitempty"`
	IsCheckBlack         bool                  `xml:"IsCheckBlack,omitempty"`
	BlackLevel           string                `xml:"BlackLevel,omitempty"`
	PixelBlackThreshold  string                `xml:"PixelBlackThreshold,omitempty"`
	SnapshotOutMode      string                `xml:"SnapshotOutMode,omitempty"`
	SpriteSnapshotConfig *SpriteSnapshotConfig `xml:"SpriteSnapshotConfig,omitempty"`
}

// AnimationVideo TODO
// 有意和转码区分，两种任务关注的参数不一样避免干扰
type AnimationVideo struct {
	Codec                      string `xml:"Codec"`
	Width                      string `xml:"Width"`
	Height                     string `xml:"Height"`
	Fps                        string `xml:"Fps"`
	AnimateOnlyKeepKeyFrame    string `xml:"AnimateOnlyKeepKeyFrame"`
	AnimateTimeIntervalOfFrame string `xml:"AnimateTimeIntervalOfFrame"`
	AnimateFramesPerSecond     string `xml:"AnimateFramesPerSecond"`
	Quality                    string `xml:"Quality"`
}

// Animation TODO
type Animation struct {
	Container    *Container      `xml:"Container,omitempty"`
	Video        *AnimationVideo `xml:"Video,omitempty"`
	TimeInterval *TimeInterval   `xml:"TimeInterval,omitempty"`
}

// HlsEncrypt TODO
type HlsEncrypt struct {
	IsHlsEncrypt bool   `xml:"IsHlsEncrypt,omitempty"`
	UriKey       string `xml:"UriKey,omitempty"`
}

// Segment TODO
type Segment struct {
	Format     string      `xml:"Format,omitempty"`
	Duration   string      `xml:"Duration,omitempty"`
	HlsEncrypt *HlsEncrypt `xml:"HlsEncrypt,omitempty"`
}

// VideoMontageVideo TODO
type VideoMontageVideo struct {
	Codec   string `xml:"Codec"`
	Width   string `xml:"Width"`
	Height  string `xml:"Height"`
	Fps     string `xml:"Fps"`
	Remove  string `xml:"Remove,omitempty"`
	Bitrate string `xml:"Bitrate"`
	Crf     string `xml:"Crf"`
}

// VideoMontage TODO
type VideoMontage struct {
	Container *Container         `xml:"Container,omitempty"`
	Video     *VideoMontageVideo `xml:"Video,omitempty"`
	Audio     *Audio             `xml:"Audio,omitempty"`
	Duration  string             `xml:"Duration,omitempty"`
}

// AudioConfig TODO
type AudioConfig struct {
	Codec      string `xml:"Codec"`
	Samplerate string `xml:"Samplerate"`
	Bitrate    string `xml:"Bitrate"`
	Channels   string `xml:"Channels"`
}

// VoiceSeparate TODO
type VoiceSeparate struct {
	AudioMode   string       `xml:"AudioMode,omitempty"` // IsAudio 人声, IsBackground 背景声, AudioAndBackground 人声和背景声
	AudioConfig *AudioConfig `xml:"AudioConfig,omitempty"`
}

// ColorEnhance TODO
type ColorEnhance struct {
	Enable     string `xml:"Enable"`
	Contrast   string `xml:"Contrast"`
	Correction string `xml:"Correction"`
	Saturation string `xml:"Saturation"`
}

// MsSharpen TODO
type MsSharpen struct {
	Enable       string `xml:"Enable"`
	SharpenLevel string `xml:"SharpenLevel"`
}

// VideoProcess TODO
type VideoProcess struct {
	ColorEnhance *ColorEnhance `xml:"ColorEnhance,omitempty"`
	MsSharpen    *MsSharpen    `xml:"MsSharpen,omitempty"`
}

// SDRtoHDR TODO
type SDRtoHDR struct {
	HdrMode string `xml:"HdrMode,omitempty"` // HLG、HDR10
}

// SuperResolution TODO
type SuperResolution struct {
	Resolution    string `xml:"Resolution,omitempty"` // sdtohd、hdto4k
	EnableScaleUp string `xml:"EnableScaleUp,omitempty"`
}

// DigitalWatermark TODO
type DigitalWatermark struct {
	Message string `xml:"Message"`
	Type    string `xml:"Type"`
	Version string `xml:"Version"`
}

// ExtractDigitalWatermark TODO
type ExtractDigitalWatermark struct {
	Type    string `xml:"Type"`
	Version string `xml:"Version"`
}

// VideoTag TODO
type VideoTag struct {
	Scenario string `xml:"Scenario"`
}

// MediaResult TODO
type MediaResult struct {
	OutputFile struct {
		Bucket  string `xml:"Bucket,omitempty"`
		Md5Info []struct {
			Md5        string `xml:"Md5,omitempty"`
			ObjectName string `xml:"ObjectName,omitempty"`
		} `xml:"Md5Info,omitempty"`
		ObjectName       []string `xml:"ObjectName,omitempty"`
		ObjectPrefix     string   `xml:"ObjectPrefix,omitempty"`
		Region           string   `xml:"Region,omitempty"`
		SpriteOutputFile struct {
			Bucket  string `xml:"Bucket,omitempty"`
			Md5Info []struct {
				Md5        string `xml:"Md5,omitempty"`
				ObjectName string `xml:"ObjectName,omitempty"`
			} `xml:"Md5Info,omitempty"`
			ObjectName   []string `xml:"ObjectName,omitempty"`
			ObjectPrefix string   `xml:"ObjectPrefix,omitempty"`
			Region       string   `xml:"Region,omitempty"`
		} `xml:"SpriteOutputFile,omitempty"`
	} `xml:"OutputFile,omitempty"`
}

// MediaInfo TODO
type MediaInfo struct {
	Format struct {
		Bitrate        string `xml:"Bitrate"`
		Duration       string `xml:"Duration"`
		FormatLongName string `xml:"FormatLongName"`
		FormatName     string `xml:"FormatName"`
		NumProgram     string `xml:"NumProgram"`
		NumStream      string `xml:"NumStream"`
		Size           string `xml:"Size"`
		StartTime      string `xml:"StartTime"`
	} `xml:"Format"`
	Stream struct {
		Audio []struct {
			Bitrate        string `xml:"Bitrate"`
			Channel        string `xml:"Channel"`
			ChannelLayout  string `xml:"ChannelLayout"`
			CodecLongName  string `xml:"CodecLongName"`
			CodecName      string `xml:"CodecName"`
			CodecTag       string `xml:"CodecTag"`
			CodecTagString string `xml:"CodecTagString"`
			CodecTimeBase  string `xml:"CodecTimeBase"`
			Duration       string `xml:"Duration"`
			Index          string `xml:"Index"`
			Language       string `xml:"Language"`
			SampleFmt      string `xml:"SampleFmt"`
			SampleRate     string `xml:"SampleRate"`
			StartTime      string `xml:"StartTime"`
			Timebase       string `xml:"Timebase"`
		} `xml:"Audio"`
		Subtitle string `xml:"Subtitle"`
		Video    []struct {
			AvgFps         string `xml:"AvgFps"`
			Bitrate        string `xml:"Bitrate"`
			CodecLongName  string `xml:"CodecLongName"`
			CodecName      string `xml:"CodecName"`
			CodecTag       string `xml:"CodecTag"`
			CodecTagString string `xml:"CodecTagString"`
			CodecTimeBase  string `xml:"CodecTimeBase"`
			Dar            string `xml:"Dar"`
			Duration       string `xml:"Duration"`
			Fps            string `xml:"Fps"`
			HasBFrame      string `xml:"HasBFrame"`
			Height         string `xml:"Height"`
			Index          string `xml:"Index"`
			Language       string `xml:"Language"`
			Level          string `xml:"Level"`
			NumFrames      string `xml:"NumFrames"`
			PixFormat      string `xml:"PixFormat"`
			Profile        string `xml:"Profile"`
			RefFrames      string `xml:"RefFrames"`
			Rotation       string `xml:"Rotation"`
			Sar            string `xml:"Sar"`
			StartTime      string `xml:"StartTime"`
			Timebase       string `xml:"Timebase"`
			Width          string `xml:"Width"`
			ColorRange     string `xml:"ColorRange"`
			ColorTransfer  string `xml:"ColorTransfer"`
			ColorPrimaries string `xml:"ColorPrimaries"`
		} `xml:"Video"`
	} `xml:"Stream"`
}

// PicProcess TODO
type PicProcess struct {
	IsPicInfo   string `xml:"IsPicInfo,omitempty"`
	ProcessRule string `xml:"ProcessRule,omitempty"`
}

// PicProcessResult TODO
type PicProcessResult struct {
	UploadResult struct {
		OriginalInfo struct {
			Key       string `xml:"Key"`
			Location  string `xml:"Location"`
			ETag      string `xml:"ETag"`
			ImageInfo struct {
				Format      string `xml:"Format"`
				Width       int32  `xml:"Width"`
				Height      int32  `xml:"Height"`
				Quality     int32  `xml:"Quality"`
				Ave         string `xml:"Ave"`
				Orientation int32  `xml:"Orientation"`
			} `xml:"ImageInfo"`
		} `xml:"OriginalInfo"`
		ProcessResults struct {
			Object struct {
				Key      string `xml:"Key"`
				Location string `xml:"Location"`
				Format   string `xml:"Format"`
				Width    int32  `xml:"Width"`
				Height   int32  `xml:"Height"`
				Size     int32  `xml:"Size"`
				Quality  int32  `xml:"Quality"`
				Etag     string `xml:"Etag"`
			} `xml:"Object"`
		} `xml:"ProcessResults"`
	} `xml:"UploadResult"`
}

// PicProcessJobOperation TODO
type PicProcessJobOperation struct {
	PicProcess       *PicProcess       `xml:"PicProcess,omitempty"`
	PicProcessResult *PicProcessResult `xml:"PicProcessResult,omitempty"`
	Output           *JobOutput        `xml:"Output,omitempty"`
}

// MediaProcessJobOperation TODO
type MediaProcessJobOperation struct {
	Tag                     string                   `xml:"Tag,omitempty"`
	Output                  *JobOutput               `xml:"Output,omitempty"`
	MediaResult             *MediaResult             `xml:"MediaResult,omitempty"`
	MediaInfo               *MediaInfo               `xml:"MediaInfo,omitempty"`
	Transcode               *Transcode               `xml:"Transcode,omitempty"`
	Watermark               []Watermark              `xml:"Watermark,omitempty"`
	TemplateId              string                   `xml:"TemplateId,omitempty"`
	WatermarkTemplateId     []string                 `xml:"WatermarkTemplateId,omitempty"`
	ConcatTemplate          *ConcatTemplate          `xml:"ConcatTemplate,omitempty"`
	Snapshot                *Snapshot                `xml:"Snapshot,omitempty"`
	Animation               *Animation               `xml:"Animation,omitempty"`
	Segment                 *Segment                 `xml:"Segment,omitempty"`
	VideoMontage            *VideoMontage            `xml:"VideoMontage,omitempty"`
	VoiceSeparate           *VoiceSeparate           `xml:"VoiceSeparate,omitempty"`
	VideoProcess            *VideoProcess            `xml:"VideoProcess,omitempty"`
	TranscodeTemplateId     string                   `xml:"TranscodeTemplateId,omitempty"` // 视频增强、超分、SDRtoHDR任务类型，可以选择转码模板相关参数
	SDRtoHDR                *SDRtoHDR                `xml:"SDRtoHDR,omitempty"`
	SuperResolution         *SuperResolution         `xml:"SuperResolution,omitempty"`
	DigitalWatermark        *DigitalWatermark        `xml:"DigitalWatermark,omitempty"`
	ExtractDigitalWatermark *ExtractDigitalWatermark `xml:"ExtractDigitalWatermark,omitempty"`
	VideoTag                *VideoTag                `xml:"VideoTag,omitempty"`
}

// CreatePicJobsOptions TODO
type CreatePicJobsOptions struct {
	XMLName   xml.Name                `xml:"Request"`
	Tag       string                  `xml:"Tag,omitempty"`
	Input     *JobInput               `xml:"Input,omitempty"`
	Operation *PicProcessJobOperation `xml:"Operation,omitempty"`
	QueueId   string                  `xml:"QueueId,omitempty"`
	CallBack  string                  `xml:"CallBack,omitempty"`
}

// CreateMediaJobsOptions TODO
type CreateMediaJobsOptions struct {
	XMLName   xml.Name                  `xml:"Request"`
	Tag       string                    `xml:"Tag,omitempty"`
	Input     *JobInput                 `xml:"Input,omitempty"`
	Operation *MediaProcessJobOperation `xml:"Operation,omitempty"`
	QueueId   string                    `xml:"QueueId,omitempty"`
	CallBack  string                    `xml:"CallBack,omitempty"`
}

// MediaProcessJobDetail TODO
type MediaProcessJobDetail struct {
	Code         string                    `xml:"Code,omitempty"`
	Message      string                    `xml:"Message,omitempty"`
	JobId        string                    `xml:"JobId,omitempty"`
	Tag          string                    `xml:"Tag,omitempty"`
	Progress     string                    `xml:"Progress,omitempty"`
	State        string                    `xml:"State,omitempty"`
	CreationTime string                    `xml:"CreationTime,omitempty"`
	QueueId      string                    `xml:"QueueId,omitempty"`
	Input        *JobInput                 `xml:"Input,omitempty"`
	Operation    *MediaProcessJobOperation `xml:"Operation,omitempty"`
}

// CreatePicJobsResult TODO
type CreatePicJobsResult CreateMediaJobsResult

// CreateMediaJobsResult TODO
type CreateMediaJobsResult struct {
	XMLName    xml.Name               `xml:"Response"`
	JobsDetail *MediaProcessJobDetail `xml:"JobsDetail,omitempty"`
}

// CreateMultiMediaJobsOptions TODO
type CreateMultiMediaJobsOptions struct {
	XMLName   xml.Name                   `xml:"Request"`
	Tag       string                     `xml:"Tag,omitempty"`
	Input     *JobInput                  `xml:"Input,omitempty"`
	Operation []MediaProcessJobOperation `xml:"Operation,omitempty"`
	QueueId   string                     `xml:"QueueId,omitempty"`
	CallBack  string                     `xml:"CallBack,omitempty"`
}

// CreateMultiMediaJobsResult TODO
type CreateMultiMediaJobsResult struct {
	XMLName    xml.Name                `xml:"Response"`
	JobsDetail []MediaProcessJobDetail `xml:"JobsDetail,omitempty"`
}

// MediaProcessJobsNotifyBody TODO
type MediaProcessJobsNotifyBody struct {
	XMLName    xml.Name `xml:"Response"`
	EventName  string   `xml:"EventName"`
	JobsDetail struct {
		Code         string `xml:"Code"`
		CreationTime string `xml:"CreationTime"`
		EndTime      string `xml:"EndTime"`
		Input        struct {
			BucketId string `xml:"BucketId"`
			Object   string `xml:"Object"`
			Region   string `xml:"Region"`
		} `xml:"Input"`
		JobId     string `xml:"JobId"`
		Message   string `xml:"Message"`
		Operation struct {
			MediaInfo   *MediaInfo   `xml:"MediaInfo"`
			MediaResult *MediaResult `xml:"MediaResult"`
			Output      struct {
				Bucket string `xml:"Bucket"`
				Object string `xml:"Object"`
				Region string `xml:"Region"`
			} `xml:"Output"`
			TemplateId   string `xml:"TemplateId"`
			TemplateName string `xml:"TemplateName"`
		} `xml:"Operation"`
		QueueId   string `xml:"QueueId"`
		StartTime string `xml:"StartTime"`
		State     string `xml:"State"`
		Tag       string `xml:"Tag"`
	} `xml:"JobsDetail"`
}

// WorkflowExecutionNotifyBody TODO
type WorkflowExecutionNotifyBody struct {
	XMLName           xml.Name `xml:"Response"`
	EventName         string   `xml:"EventName"`
	WorkflowExecution struct {
		RunId      string `xml:"RunId"`
		BucketId   string `xml:"BucketId"`
		Object     string `xml:"Object"`
		CosHeaders []struct {
			Key   string `xml:"Key"`
			Value string `xml:"Value"`
		} `xml:"CosHeaders"`
		WorkflowId   string `xml:"WorkflowId"`
		WorkflowName string `xml:"WorkflowName"`
		CreateTime   string `xml:"CreateTime"`
		State        string `xml:"State"`
		Tasks        []struct {
			Type                  string `xml:"Type"`
			CreateTime            string `xml:"CreateTime"`
			EndTime               string `xml:"EndTime"`
			State                 string `xml:"State"`
			JobId                 string `xml:"JobId"`
			Name                  string `xml:"Name"`
			TemplateId            string `xml:"TemplateId"`
			TemplateName          string `xml:"TemplateName"`
			TranscodeTemplateId   string `xml:"TranscodeTemplateId,omitempty"`
			TranscodeTemplateName string `xml:"TranscodeTemplateName,omitempty"`
			HdrMode               string `xml:"HdrMode,omitempty"`
		} `xml:"Tasks"`
	} `xml:"WorkflowExecution"`
}

// CreateMultiMediaJobs TODO
func (s *CIService) CreateMultiMediaJobs(ctx context.Context, opt *CreateMultiMediaJobsOptions) (*CreateMultiMediaJobsResult, *Response, error) {
	var res CreateMultiMediaJobsResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/jobs",
		method:  http.MethodPost,
		body:    opt,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// CreateMediaJobs TODO
func (s *CIService) CreateMediaJobs(ctx context.Context, opt *CreateMediaJobsOptions) (*CreateMediaJobsResult, *Response, error) {
	var res CreateMediaJobsResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/jobs",
		method:  http.MethodPost,
		body:    opt,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// CreatePicProcessJobs TODO
func (s *CIService) CreatePicProcessJobs(ctx context.Context, opt *CreatePicJobsOptions) (*CreatePicJobsResult, *Response, error) {
	var res CreatePicJobsResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/pic_jobs",
		method:  http.MethodPost,
		body:    opt,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// DescribePicProcessJobResult TODO
type DescribePicProcessJobResult DescribeMediaProcessJobResult

// DescribeMediaProcessJobResult TODO
type DescribeMediaProcessJobResult struct {
	XMLName        xml.Name               `xml:"Response"`
	JobsDetail     *MediaProcessJobDetail `xml:"JobsDetail,omitempty"`
	NonExistJobIds string                 `xml:"NonExistJobIds,omitempty"`
}

// DescribeMediaJob TODO
func (s *CIService) DescribeMediaJob(ctx context.Context, jobid string) (*DescribeMediaProcessJobResult, *Response, error) {
	var res DescribeMediaProcessJobResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/jobs/" + jobid,
		method:  http.MethodGet,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// DescribePicProcessJob TODO
func (s *CIService) DescribePicProcessJob(ctx context.Context, jobid string) (*DescribePicProcessJobResult, *Response, error) {
	var res DescribePicProcessJobResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/pic_jobs/" + jobid,
		method:  http.MethodGet,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// DescribeMutilMediaProcessJobResult TODO
type DescribeMutilMediaProcessJobResult struct {
	XMLName        xml.Name                `xml:"Response"`
	JobsDetail     []MediaProcessJobDetail `xml:"JobsDetail,omitempty"`
	NonExistJobIds []string                `xml:"NonExistJobIds,omitempty"`
}

// DescribeMultiMediaJob TODO
func (s *CIService) DescribeMultiMediaJob(ctx context.Context, jobids []string) (*DescribeMutilMediaProcessJobResult, *Response, error) {
	jobidsStr := ""
	if len(jobids) < 1 {
		return nil, nil, errors.New("empty param jobids")
	} else {
		jobidsStr = strings.Join(jobids, ",")
	}

	var res DescribeMutilMediaProcessJobResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/jobs/" + jobidsStr,
		method:  http.MethodGet,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// DescribeMediaJobsOptions TODO
type DescribeMediaJobsOptions struct {
	QueueId           string `url:"queueId,omitempty"`
	Tag               string `url:"tag,omitempty"`
	OrderByTime       string `url:"orderByTime,omitempty"`
	NextToken         string `url:"nextToken,omitempty"`
	Size              int    `url:"size,omitempty"`
	States            string `url:"states,omitempty"`
	StartCreationTime string `url:"startCreationTime,omitempty"`
	EndCreationTime   string `url:"endCreationTime,omitempty"`
}

// DescribeMediaJobsResult TODO
type DescribeMediaJobsResult struct {
	XMLName    xml.Name                `xml:"Response"`
	JobsDetail []MediaProcessJobDetail `xml:"JobsDetail,omitempty"`
	NextToken  string                  `xml:"NextToken,omitempty"`
}

// DescribeMediaJobs TODO
// https://cloud.tencent.com/document/product/460/48235
func (s *CIService) DescribeMediaJobs(ctx context.Context, opt *DescribeMediaJobsOptions) (*DescribeMediaJobsResult, *Response, error) {
	var res DescribeMediaJobsResult
	sendOpt := sendOptions{
		baseURL:  s.client.BaseURL.CIURL,
		uri:      "/jobs",
		optQuery: opt,
		method:   http.MethodGet,
		result:   &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// DescribePicProcessQueuesOptions TODO
type DescribePicProcessQueuesOptions DescribeMediaProcessQueuesOptions

// DescribeMediaProcessQueuesOptions TODO
type DescribeMediaProcessQueuesOptions struct {
	QueueIds   string `url:"queueIds,omitempty"`
	State      string `url:"state,omitempty"`
	PageNumber int    `url:"pageNumber,omitempty"`
	PageSize   int    `url:"pageSize,omitempty"`
}

// DescribePicProcessQueuesResult TODO
type DescribePicProcessQueuesResult DescribeMediaProcessQueuesResult

// DescribeMediaProcessQueuesResult TODO
type DescribeMediaProcessQueuesResult struct {
	XMLName      xml.Name            `xml:"Response"`
	RequestId    string              `xml:"RequestId,omitempty"`
	TotalCount   int                 `xml:"TotalCount,omitempty"`
	PageNumber   int                 `xml:"PageNumber,omitempty"`
	PageSize     int                 `xml:"PageSize,omitempty"`
	QueueList    []MediaProcessQueue `xml:"QueueList,omitempty"`
	NonExistPIDs []string            `xml:"NonExistPIDs,omitempty"`
}

// MediaProcessQueue TODO
type MediaProcessQueue struct {
	QueueId       string                         `xml:"QueueId,omitempty"`
	Name          string                         `xml:"Name,omitempty"`
	State         string                         `xml:"State,omitempty"`
	MaxSize       int                            `xml:"MaxSize,omitempty"`
	MaxConcurrent int                            `xml:"MaxConcurrent,omitempty"`
	UpdateTime    string                         `xml:"UpdateTime,omitempty"`
	CreateTime    string                         `xml:"CreateTime,omitempty"`
	NotifyConfig  *MediaProcessQueueNotifyConfig `xml:"NotifyConfig,omitempty"`
}

// MediaProcessQueueNotifyConfig TODO
type MediaProcessQueueNotifyConfig struct {
	Url   string `xml:"Url,omitempty"`
	State string `xml:"State,omitempty"`
	Type  string `xml:"Type,omitempty"`
	Event string `xml:"Event,omitempty"`
}

// DescribeMediaProcessQueues TODO
func (s *CIService) DescribeMediaProcessQueues(ctx context.Context, opt *DescribeMediaProcessQueuesOptions) (*DescribeMediaProcessQueuesResult, *Response, error) {
	var res DescribeMediaProcessQueuesResult
	sendOpt := sendOptions{
		baseURL:  s.client.BaseURL.CIURL,
		uri:      "/queue",
		optQuery: opt,
		method:   http.MethodGet,
		result:   &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// DescribePicProcessQueues TODO
func (s *CIService) DescribePicProcessQueues(ctx context.Context, opt *DescribePicProcessQueuesOptions) (*DescribePicProcessQueuesResult, *Response, error) {
	var res DescribePicProcessQueuesResult
	sendOpt := sendOptions{
		baseURL:  s.client.BaseURL.CIURL,
		uri:      "/picqueue",
		optQuery: opt,
		method:   http.MethodGet,
		result:   &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// UpdateMediaProcessQueueOptions TODO
type UpdateMediaProcessQueueOptions struct {
	XMLName      xml.Name                       `xml:"Request"`
	Name         string                         `xml:"Name,omitempty"`
	QueueID      string                         `xml:"QueueID,omitempty"`
	State        string                         `xml:"State,omitempty"`
	NotifyConfig *MediaProcessQueueNotifyConfig `xml:"NotifyConfig,omitempty"`
}

// UpdateMediaProcessQueueResult TODO
type UpdateMediaProcessQueueResult struct {
	XMLName   xml.Name           `xml:"Response"`
	RequestId string             `xml:"RequestId"`
	Queue     *MediaProcessQueue `xml:"Queue"`
}

// UpdateMediaProcessQueue TODO
func (s *CIService) UpdateMediaProcessQueue(ctx context.Context, opt *UpdateMediaProcessQueueOptions) (*UpdateMediaProcessQueueResult, *Response, error) {
	var res UpdateMediaProcessQueueResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/queue/" + opt.QueueID,
		body:    opt,
		method:  http.MethodPut,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// DescribeMediaProcessBucketsOptions TODO
type DescribeMediaProcessBucketsOptions struct {
	Regions     string `url:"regions,omitempty"`
	BucketNames string `url:"bucketNames,omitempty"`
	BucketName  string `url:"bucketName,omitempty"`
	PageNumber  int    `url:"pageNumber,omitempty"`
	PageSize    int    `url:"pageSize,omitempty"`
}

// DescribeMediaProcessBucketsResult TODO
type DescribeMediaProcessBucketsResult struct {
	XMLName         xml.Name             `xml:"Response"`
	RequestId       string               `xml:"RequestId,omitempty"`
	TotalCount      int                  `xml:"TotalCount,omitempty"`
	PageNumber      int                  `xml:"PageNumber,omitempty"`
	PageSize        int                  `xml:"PageSize,omitempty"`
	MediaBucketList []MediaProcessBucket `xml:"MediaBucketList,omitempty"`
}

// MediaProcessBucket TODO
type MediaProcessBucket struct {
	BucketId   string `xml:"BucketId,omitempty"`
	Region     string `xml:"Region,omitempty"`
	CreateTime string `xml:"CreateTime,omitempty"`
}

// DescribeMediaProcessBuckets TODO
// 媒体bucket接口 https://cloud.tencent.com/document/product/436/48988
func (s *CIService) DescribeMediaProcessBuckets(ctx context.Context, opt *DescribeMediaProcessBucketsOptions) (*DescribeMediaProcessBucketsResult, *Response, error) {
	var res DescribeMediaProcessBucketsResult
	sendOpt := sendOptions{
		baseURL:  s.client.BaseURL.CIURL,
		uri:      "/mediabucket",
		optQuery: opt,
		method:   http.MethodGet,
		result:   &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// GetMediaInfoResult TODO
type GetMediaInfoResult struct {
	XMLName   xml.Name   `xml:"Response"`
	MediaInfo *MediaInfo `xml:"MediaInfo"`
}

// GetMediaInfo TODO
// 媒体信息接口 https://cloud.tencent.com/document/product/436/55672
func (s *CIService) GetMediaInfo(ctx context.Context, name string, opt *ObjectGetOptions, id ...string) (*GetMediaInfoResult, *Response, error) {
	var u string
	if len(id) == 1 {
		u = fmt.Sprintf("/%s?versionId=%s&ci-process=videoinfo", encodeURIComponent(name), id[0])
	} else if len(id) == 0 {
		u = fmt.Sprintf("/%s?ci-process=videoinfo", encodeURIComponent(name))
	} else {
		return nil, nil, fmt.Errorf("wrong params")
	}

	var res GetMediaInfoResult
	sendOpt := sendOptions{
		baseURL:   s.client.BaseURL.BucketURL,
		uri:       u,
		method:    http.MethodGet,
		optQuery:  opt,
		optHeader: opt,
		result:    &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// GenerateMediaInfoOptions TODO
type GenerateMediaInfoOptions struct {
	XMLName xml.Name  `xml:"Request"`
	Input   *JobInput `xml:"Input,omitempty"`
}

// GenerateMediaInfo TODO
// 生成媒体信息接口，支持大文件，耗时较大请求
func (s *CIService) GenerateMediaInfo(ctx context.Context, opt *GenerateMediaInfoOptions) (*GetMediaInfoResult, *Response, error) {

	var res GetMediaInfoResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/mediainfo",
		method:  http.MethodPost,
		body:    opt,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// GetSnapshotOptions TODO
type GetSnapshotOptions struct {
	Time   float32 `url:"time,omitempty"`
	Height int     `url:"height,omitempty"`
	Width  int     `url:"width,omitempty"`
	Format string  `url:"format,omitempty"`
	Rotate string  `url:"rotate,omitempty"`
	Mode   string  `url:"mode,omitempty"`
}

// GetSnapshot TODO
// 媒体截图接口 https://cloud.tencent.com/document/product/436/55671
func (s *CIService) GetSnapshot(ctx context.Context, name string, opt *GetSnapshotOptions, id ...string) (*Response, error) {
	var u string
	if len(id) == 1 {
		u = fmt.Sprintf("/%s?versionId=%s&ci-process=snapshot", encodeURIComponent(name), id[0])
	} else if len(id) == 0 {
		u = fmt.Sprintf("/%s?ci-process=snapshot", encodeURIComponent(name))
	} else {
		return nil, fmt.Errorf("wrong params")
	}

	sendOpt := sendOptions{
		baseURL:          s.client.BaseURL.BucketURL,
		uri:              u,
		method:           http.MethodGet,
		optQuery:         opt,
		disableCloseBody: true,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return resp, err
}

// GetPrivateM3U8Options TODO
type GetPrivateM3U8Options struct {
	Expires int `url:"expires"`
}

// GetPrivateM3U8 TODO
// 获取私有m3u8资源接口 https://cloud.tencent.com/document/product/460/63738
func (s *CIService) GetPrivateM3U8(ctx context.Context, name string, opt *GetPrivateM3U8Options, id ...string) (*Response, error) {
	var u string
	if len(id) == 1 {
		u = fmt.Sprintf("/%s?versionId=%s&ci-process=pm3u8", encodeURIComponent(name), id[0])
	} else if len(id) == 0 {
		u = fmt.Sprintf("/%s?ci-process=pm3u8", encodeURIComponent(name))
	} else {
		return nil, fmt.Errorf("wrong params")
	}

	sendOpt := sendOptions{
		baseURL:          s.client.BaseURL.BucketURL,
		uri:              u,
		method:           http.MethodGet,
		optQuery:         opt,
		disableCloseBody: true,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return resp, err
}

// TriggerWorkflowOptions TODO
type TriggerWorkflowOptions struct {
	WorkflowId string `url:"workflowId"`
	Object     string `url:"object"`
}

// TriggerWorkflowResult TODO
type TriggerWorkflowResult struct {
	XMLName    xml.Name `xml:"Response"`
	InstanceId string   `xml:"InstanceId"`
	RequestId  string   `xml:"RequestId"`
}

// TriggerWorkflow TODO
// 单文件触发工作流 https://cloud.tencent.com/document/product/460/54640
func (s *CIService) TriggerWorkflow(ctx context.Context, opt *TriggerWorkflowOptions) (*TriggerWorkflowResult, *Response, error) {
	var res TriggerWorkflowResult
	sendOpt := sendOptions{
		baseURL:  s.client.BaseURL.CIURL,
		uri:      "/triggerworkflow",
		optQuery: opt,
		method:   http.MethodPost,
		result:   &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// DescribeWorkflowExecutionsOptions TODO
type DescribeWorkflowExecutionsOptions struct {
	WorkflowId        string `url:"workflowId,omitempty"`
	Name              string `url:"Name,omitempty"`
	OrderByTime       string `url:"orderByTime,omitempty"`
	NextToken         string `url:"nextToken,omitempty"`
	Size              int    `url:"size,omitempty"`
	States            string `url:"states,omitempty"`
	StartCreationTime string `url:"startCreationTime,omitempty"`
	EndCreationTime   string `url:"endCreationTime,omitempty"`
}

// WorkflowExecutionList TODO
type WorkflowExecutionList struct {
	RunId        string `xml:"RunId,omitempty"`
	WorkflowId   string `xml:"WorkflowId,omitempty"`
	State        string `xml:"State,omitempty"`
	CreationTime string `xml:"CreationTime,omitempty"`
	Object       string `xml:"Object,omitempty"`
}

// DescribeWorkflowExecutionsResult TODO
type DescribeWorkflowExecutionsResult struct {
	XMLName               xml.Name                `xml:"Response"`
	WorkflowExecutionList []WorkflowExecutionList `xml:"WorkflowExecutionList,omitempty"`
	NextToken             string                  `xml:"NextToken,omitempty"`
}

// DescribeWorkflowExecutions TODO
// 获取工作流实例列表 https://cloud.tencent.com/document/product/460/45950
func (s *CIService) DescribeWorkflowExecutions(ctx context.Context, opt *DescribeWorkflowExecutionsOptions) (*DescribeWorkflowExecutionsResult, *Response, error) {
	var res DescribeWorkflowExecutionsResult
	sendOpt := sendOptions{
		baseURL:  s.client.BaseURL.CIURL,
		uri:      "/workflowexecution",
		optQuery: opt,
		method:   http.MethodGet,
		result:   &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// NotifyConfig TODO
type NotifyConfig struct {
	URL   string `xml:"Url,omitempty"`
	Event string `xml:"Event,omitempty"`
	Type  string `xml:"Type,omitempty"`
}

// ExtFilter TODO
type ExtFilter struct {
	State      string `xml:"State,omitempty"`
	Audio      string `xml:"Audio,omitempty"`
	Custom     string `xml:"Custom,omitempty"`
	CustomExts string `xml:"CustomExts,omitempty"`
	AllFile    string `xml:"AllFile,omitempty"`
}

// NodeInput TODO
type NodeInput struct {
	QueueId      string        `xml:"QueueId,omitempty"`
	ObjectPrefix string        `xml:"ObjectPrefix,omitempty"`
	NotifyConfig *NotifyConfig `xml:"NotifyConfig,omitempty" json:"NotifyConfig,omitempty"`
	ExtFilter    *ExtFilter    `xml:"ExtFilter,omitempty" json:"ExtFilter,omitempty"`
}

// NodeSegment TODO
type NodeSegment struct {
	Format   string `xml:"Format,omitempty"`
	Duration string `xml:"Duration,omitempty"`
}

// NodeOutput TODO
type NodeOutput struct {
	Region       string `xml:"Region,omitempty"`
	Bucket       string `xml:"Bucket,omitempty"`
	Object       string `xml:"Object,omitempty"`
	AuObject     string `xml:"AuObject,omitempty"`
	SpriteObject string `xml:"SpriteObject,omitempty"`
}

// NodeSCF TODO
type NodeSCF struct {
	Region       string `xml:"Region,omitempty"`
	FunctionName string `xml:"FunctionName,omitempty"`
	Namespace    string `xml:"Namespace,omitempty"`
}

// NodeSDRtoHDR TODO
type NodeSDRtoHDR struct {
	HdrMode string `xml:"HdrMode,omitempty"`
}

// NodeSmartCover TODO
type NodeSmartCover struct {
	Format           string `xml:"Format,omitempty"`
	Width            string `xml:"Width,omitempty"`
	Height           string `xml:"Height,omitempty"`
	Count            string `xml:"Count,omitempty"`
	DeleteDuplicates string `xml:"DeleteDuplicates,omitempty"`
}

// NodeOperation TODO
type NodeOperation struct {
	TemplateId          string          `xml:"TemplateId,omitempty" json:"TemplateId,omitempty"`
	Segment             *NodeSegment    `xml:"Segment,omitempty" json:"Segment,omitempty" `
	Output              *NodeOutput     `xml:"Output,omitempty" json:"Output,omitempty"`
	SCF                 *NodeSCF        `xml:"SCF,omitempty" json:"SCF,omitempty"`
	SDRtoHDR            *NodeSDRtoHDR   `xml:"SDRtoHDR,omitempty" json:"SDRtoHDR,omitempty"`
	SmartCover          *NodeSmartCover `xml:"SmartCover,omitempty" json:"SmartCover,omitempty"`
	WatermarkTemplateId string          `xml:"WatermarkTemplateId,omitempty" json:"WatermarkTemplateId,omitempty`
	TranscodeTemplateId string          `xml:"TranscodeTemplateId,omitempty" json:"TranscodeTemplateId,omitempty"`
}

// Node TODO
type Node struct {
	Type      string         `xml:"Type"`
	Input     *NodeInput     `xml:"Input,omitempty" json:"Input,omitempty"`
	Operation *NodeOperation `xml:"Operation,omitempty" json:"Operation,omitempty"`
}

// Topology TODO
type Topology struct {
	Dependencies map[string]string `json:"Dependencies,omitempty"`
	Nodes        map[string]Node   `json:"Nodes,omitempty"`
}

// UnmarshalXML TODO
func (m *Topology) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v struct {
		XMLName      xml.Name //`xml:"Topology"`
		Dependencies struct {
			Inner []byte `xml:",innerxml"`
		} `xml:"Dependencies"`
		Nodes struct {
			Inner []byte `xml:",innerxml"`
		} `xml:"Nodes"`
	}
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}

	myMap := make(map[string]interface{})

	// ... do the mxj magic here ... -

	temp := v.Nodes.Inner

	prefix := "<Nodes>"
	postfix := "</Nodes>"
	str := prefix + string(temp) + postfix
	//fmt.Println(str)
	myMxjMap, _ := mxj.NewMapXml([]byte(str))
	myMap, _ = myMxjMap["Nodes"].(map[string]interface{})
	nodesMap := make(map[string]Node)

	for k, v := range myMap {
		var node Node
		mapstructure.Decode(v, &node)
		nodesMap[k] = node
	}

	// fill myMap
	m.Nodes = nodesMap

	deps := make(map[string]interface{})

	tep := "<Dependencies>" + string(v.Dependencies.Inner) + "</Dependencies>"
	tepMxjMap, _ := mxj.NewMapXml([]byte(tep))
	deps, _ = tepMxjMap["Dependencies"].(map[string]interface{})
	depsString := make(map[string]string)
	for k, v := range deps {
		depsString[k] = v.(string)
	}
	m.Dependencies = depsString
	return nil
}

// WorkflowExecution TODO
type WorkflowExecution struct {
	RunId        string   `xml:"RunId,omitempty" json:"RunId,omitempty"`
	WorkflowId   string   `xml:"WorkflowId,omitempty" json:"WorkflowId,omitempty"`
	WorkflowName string   `xml:"WorkflowName,omitempty" json:"WorkflowName,omitempty"`
	State        string   `xml:"State,omitempty" json:"State,omitempty"`
	CreateTime   string   `xml:"CreateTime,omitempty" json:"CreateTime,omitempty"`
	Object       string   `xml:"Object,omitempty" json:"Object,omitempty"`
	Topology     Topology `xml:"Topology,omitempty" json:"Topology,omitempty"`
}

// DescribeWorkflowExecutionResult TODO
type DescribeWorkflowExecutionResult struct {
	XMLName           xml.Name            `xml:"Response"`
	WorkflowExecution []WorkflowExecution `xml:"WorkflowExecution,omitempty" json:"WorkflowExecution,omitempty"`
	NextToken         string              `xml:"NextToken,omitempty" json:"NextToken,omitempty"`
}

// DescribeWorkflowExecution TODO
// 获取工作流实例详情 https://cloud.tencent.com/document/product/460/45949
func (s *CIService) DescribeWorkflowExecution(ctx context.Context, runId string) (*DescribeWorkflowExecutionResult, *Response, error) {
	var res DescribeWorkflowExecutionResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/workflowexecution/" + runId,
		method:  http.MethodGet,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// SpeechRecognition TODO
type SpeechRecognition struct {
	ChannelNum      string `xml:"ChannelNum,omitempty"`
	ConvertNumMode  string `xml:"ConvertNumMode,omitempty"`
	EngineModelType string `xml:"EngineModelType,omitempty"`
	FilterDirty     string `xml:"FilterDirty,omitempty"`
	FilterModal     string `xml:"FilterModal,omitempty"`
	ResTextFormat   string `xml:"ResTextFormat,omitempty"`
}

// SpeechRecognitionResult TODO
type SpeechRecognitionResult struct {
	AudioTime float64  `xml:"AudioTime,omitempty"`
	Result    []string `xml:"Result,omitempty"`
}

// ASRJobOperation TODO
type ASRJobOperation struct {
	Tag                     string                   `xml:"Tag,omitempty"`
	Output                  *JobOutput               `xml:"Output,omitempty"`
	SpeechRecognition       *SpeechRecognition       `xml:"SpeechRecognition,omitempty"`
	SpeechRecognitionResult *SpeechRecognitionResult `xml:"SpeechRecognitionResult,omitempty"`
}

// CreateASRJobsOptions TODO
type CreateASRJobsOptions struct {
	XMLName   xml.Name         `xml:"Request"`
	Tag       string           `xml:"Tag,omitempty"`
	Input     *JobInput        `xml:"Input,omitempty"`
	Operation *ASRJobOperation `xml:"Operation,omitempty"`
	QueueId   string           `xml:"QueueId,omitempty"`
	CallBack  string           `xml:"CallBack,omitempty"`
}

// ASRJobDetail TODO
type ASRJobDetail struct {
	Code         string           `xml:"Code,omitempty"`
	Message      string           `xml:"Message,omitempty"`
	JobId        string           `xml:"JobId,omitempty"`
	Tag          string           `xml:"Tag,omitempty"`
	State        string           `xml:"State,omitempty"`
	CreationTime string           `xml:"CreationTime,omitempty"`
	QueueId      string           `xml:"QueueId,omitempty"`
	Input        *JobInput        `xml:"Input,omitempty"`
	Operation    *ASRJobOperation `xml:"Operation,omitempty"`
}

// CreateASRJobsResult TODO
type CreateASRJobsResult struct {
	XMLName    xml.Name      `xml:"Response"`
	JobsDetail *ASRJobDetail `xml:"JobsDetail,omitempty"`
}

// CreateASRJobs TODO
func (s *CIService) CreateASRJobs(ctx context.Context, opt *CreateASRJobsOptions) (*CreateASRJobsResult, *Response, error) {
	var res CreateASRJobsResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/asr_jobs",
		method:  http.MethodPost,
		body:    opt,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// DescribeMutilASRJobResult TODO
type DescribeMutilASRJobResult struct {
	XMLName        xml.Name       `xml:"Response"`
	JobsDetail     []ASRJobDetail `xml:"JobsDetail,omitempty"`
	NonExistJobIds []string       `xml:"NonExistJobIds,omitempty"`
}

// DescribeMultiASRJob TODO
func (s *CIService) DescribeMultiASRJob(ctx context.Context, jobids []string) (*DescribeMutilASRJobResult, *Response, error) {
	jobidsStr := ""
	if len(jobids) < 1 {
		return nil, nil, errors.New("empty param jobids")
	} else {
		jobidsStr = strings.Join(jobids, ",")
	}

	var res DescribeMutilASRJobResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/asr_jobs/" + jobidsStr,
		method:  http.MethodGet,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// DescribeMediaTemplateOptions TODO
type DescribeMediaTemplateOptions struct {
	Tag        string `url:"tag,omitempty"`
	Category   string `url:"category,omitempty"`
	Ids        string `url:"ids,omitempty"`
	Name       string `url:"name,omitempty"`
	PageNumber int    `url:"pageNumber,omitempty"`
	PageSize   int    `url:"pageSize,omitempty"`
}

// DescribeMediaTemplateResult TODO
type DescribeMediaTemplateResult struct {
	XMLName      xml.Name   `xml:"Response"`
	TemplateList []Template `xml:"TemplateList,omitempty"`
	RequestId    string     `xml:"RequestId,omitempty"`
	TotalCount   int        `xml:"TotalCount,omitempty"`
	PageNumber   int        `xml:"PageNumber,omitempty"`
	PageSize     int        `xml:"PageSize,omitempty"`
}

// DescribeMediaTemplate 搜索模板
func (s *CIService) DescribeMediaTemplate(ctx context.Context, opt *DescribeMediaTemplateOptions) (*DescribeMediaTemplateResult, *Response, error) {
	var res DescribeMediaTemplateResult
	sendOpt := sendOptions{
		baseURL:  s.client.BaseURL.CIURL,
		uri:      "/template",
		optQuery: opt,
		method:   http.MethodGet,
		result:   &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// DeleteMediaTemplateResult TODO
type DeleteMediaTemplateResult struct {
	RequestId  string `xml:"RequestId,omitempty"`
	TemplateId string `xml:"TemplateId,omitempty"`
}

// DeleteMediaTemplate TODO
func (s *CIService) DeleteMediaTemplate(ctx context.Context, tempalteId string) (*DeleteMediaTemplateResult, *Response, error) {
	var res DeleteMediaTemplateResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/template/" + tempalteId,
		method:  http.MethodDelete,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// CreateMediaSnapshotTemplateOptions TODO
type CreateMediaSnapshotTemplateOptions struct {
	XMLName  xml.Name  `xml:"Request"`
	Tag      string    `xml:"Tag,omitempty"`
	Name     string    `xml:"Name,omitempty"`
	Snapshot *Snapshot `xml:"Snapshot,omitempty"`
}

// CreateMediaTranscodeTemplateOptions TODO
type CreateMediaTranscodeTemplateOptions struct {
	XMLName      xml.Name      `xml:"Request"`
	Tag          string        `xml:"Tag,omitempty"`
	Name         string        `xml:"Name,omitempty"`
	Container    *Container    `xml:"Container,omitempty"`
	Video        *Video        `xml:"Video,omitempty"`
	Audio        *Audio        `xml:"Audio,omitempty"`
	TimeInterval *TimeInterval `xml:"TimeInterval,omitempty"`
	TransConfig  *TransConfig  `xml:"TransConfig,omitempty"`
}

// CreateMediaAnimationTemplateOptions TODO
type CreateMediaAnimationTemplateOptions struct {
	XMLName      xml.Name        `xml:"Request"`
	Tag          string          `xml:"Tag,omitempty"`
	Name         string          `xml:"Name,omitempty"`
	Container    *Container      `xml:"Container,omitempty"`
	Video        *AnimationVideo `xml:"Video,omitempty"`
	TimeInterval *TimeInterval   `xml:"TimeInterval,omitempty"`
}

// CreateMediaConcatTemplateOptions TODO
type CreateMediaConcatTemplateOptions struct {
	XMLName        xml.Name        `xml:"Request"`
	Tag            string          `xml:"Tag,omitempty"`
	Name           string          `xml:"Name,omitempty"`
	ConcatTemplate *ConcatTemplate `xml:"ConcatTemplate,omitempty"`
}

// CreateMediaVideoProcessTemplateOptions TODO
type CreateMediaVideoProcessTemplateOptions struct {
	XMLName      xml.Name      `xml:"Request"`
	Tag          string        `xml:"Tag,omitempty"`
	Name         string        `xml:"Name,omitempty"`
	ColorEnhance *ColorEnhance `xml:"ColorEnhance,omitempty"`
	MsSharpen    *MsSharpen    `xml:"MsSharpen,omitempty"`
}

// CreateMediaTemplateResult TODO
type CreateMediaTemplateResult struct {
	XMLName   xml.Name  `xml:"Response"`
	RequestId string    `xml:"RequestId,omitempty"`
	Template  *Template `xml:"Template,omitempty"`
}

// Template TODO
type Template struct {
	TemplateId     string          `xml:"TemplateId,omitempty"`
	Tag            string          `xml:"Code,omitempty"`
	Name           string          `xml:"Name,omitempty"`
	TransTpl       *Transcode      `xml:"TransTpl,omitempty"`
	CreationTime   string          `xml:"CreationTime,omitempty"`
	UpdateTime     string          `xml:"UpdateTime,omitempty"`
	BucketId       string          `xml:"BucketId,omitempty"`
	Category       string          `xml:"Category,omitempty"`
	Snapshot       *Snapshot       `xml:"Snapshot,omitempty"`
	Animation      *Animation      `xml:"Animation,omitempty"`
	ConcatTemplate *ConcatTemplate `xml:"ConcatTemplate,omitempty"`
	VideoProcess   *VideoProcess   `xml:"VideoProcess,omitempty"`
}

// CreateMediaSnapshotTemplate 创建截图模板
func (s *CIService) CreateMediaSnapshotTemplate(ctx context.Context, opt *CreateMediaSnapshotTemplateOptions) (*CreateMediaTemplateResult, *Response, error) {
	var res CreateMediaTemplateResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/template",
		method:  http.MethodPost,
		body:    opt,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// UpdateMediaSnapshotTemplate 更新截图模板
func (s *CIService) UpdateMediaSnapshotTemplate(ctx context.Context, opt *CreateMediaSnapshotTemplateOptions, templateId string) (*CreateMediaTemplateResult, *Response, error) {
	var res CreateMediaTemplateResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/template/" + templateId,
		method:  http.MethodPut,
		body:    opt,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// CreateMediaTranscodeTemplate Options 创建转码模板
func (s *CIService) CreateMediaTranscodeTemplate(ctx context.Context, opt *CreateMediaTranscodeTemplateOptions) (*CreateMediaTemplateResult, *Response, error) {
	var res CreateMediaTemplateResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/template",
		method:  http.MethodPost,
		body:    opt,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// UpdateMediaTranscodeTemplate 更新转码模板
func (s *CIService) UpdateMediaTranscodeTemplate(ctx context.Context, opt *CreateMediaTranscodeTemplateOptions, templateId string) (*CreateMediaTemplateResult, *Response, error) {
	var res CreateMediaTemplateResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/template/" + templateId,
		method:  http.MethodPut,
		body:    opt,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// CreateMediaAnimationTemplate 创建动图模板
func (s *CIService) CreateMediaAnimationTemplate(ctx context.Context, opt *CreateMediaAnimationTemplateOptions) (*CreateMediaTemplateResult, *Response, error) {
	var res CreateMediaTemplateResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/template",
		method:  http.MethodPost,
		body:    opt,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// UpdateMediaAnimationTemplate 创建动图模板
func (s *CIService) UpdateMediaAnimationTemplate(ctx context.Context, opt *CreateMediaAnimationTemplateOptions, templateId string) (*CreateMediaTemplateResult, *Response, error) {
	var res CreateMediaTemplateResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/template/" + templateId,
		method:  http.MethodPut,
		body:    opt,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// CreateMediaConcatTemplate 创建拼接模板
func (s *CIService) CreateMediaConcatTemplate(ctx context.Context, opt *CreateMediaConcatTemplateOptions) (*CreateMediaTemplateResult, *Response, error) {
	var res CreateMediaTemplateResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/template",
		method:  http.MethodPost,
		body:    opt,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// UpdateMediaConcatTemplate 创建拼接模板
func (s *CIService) UpdateMediaConcatTemplate(ctx context.Context, opt *CreateMediaConcatTemplateOptions, templateId string) (*CreateMediaTemplateResult, *Response, error) {
	var res CreateMediaTemplateResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/template/" + templateId,
		method:  http.MethodPut,
		body:    opt,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// CreateMediaVideoProcessTemplate 创建视频增强模板
func (s *CIService) CreateMediaVideoProcessTemplate(ctx context.Context, opt *CreateMediaVideoProcessTemplateOptions) (*CreateMediaTemplateResult, *Response, error) {
	var res CreateMediaTemplateResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/template",
		method:  http.MethodPost,
		body:    opt,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// UpdateMediaVideoProcessTemplate 创建视频增强模板
func (s *CIService) UpdateMediaVideoProcessTemplate(ctx context.Context, opt *CreateMediaVideoProcessTemplateOptions, templateId string) (*CreateMediaTemplateResult, *Response, error) {
	var res CreateMediaTemplateResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/template/" + templateId,
		method:  http.MethodPut,
		body:    opt,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}
