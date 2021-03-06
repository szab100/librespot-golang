// Code generated by protoc-gen-go. DO NOT EDIT.
// source: radio.proto

package Spotify

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RadioRequest struct {
	Uris             []string `protobuf:"bytes,1,rep,name=uris" json:"uris,omitempty"`
	Salt             *int32   `protobuf:"varint,2,opt,name=salt" json:"salt,omitempty"`
	Length           *int32   `protobuf:"varint,4,opt,name=length" json:"length,omitempty"`
	StationId        *string  `protobuf:"bytes,5,opt,name=stationId" json:"stationId,omitempty"`
	LastTracks       []string `protobuf:"bytes,6,rep,name=lastTracks" json:"lastTracks,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *RadioRequest) Reset()                    { *m = RadioRequest{} }
func (m *RadioRequest) String() string            { return proto.CompactTextString(m) }
func (*RadioRequest) ProtoMessage()               {}
func (*RadioRequest) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

func (m *RadioRequest) GetUris() []string {
	if m != nil {
		return m.Uris
	}
	return nil
}

func (m *RadioRequest) GetSalt() int32 {
	if m != nil && m.Salt != nil {
		return *m.Salt
	}
	return 0
}

func (m *RadioRequest) GetLength() int32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *RadioRequest) GetStationId() string {
	if m != nil && m.StationId != nil {
		return *m.StationId
	}
	return ""
}

func (m *RadioRequest) GetLastTracks() []string {
	if m != nil {
		return m.LastTracks
	}
	return nil
}

type MultiSeedRequest struct {
	Uris             []string `protobuf:"bytes,1,rep,name=uris" json:"uris,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MultiSeedRequest) Reset()                    { *m = MultiSeedRequest{} }
func (m *MultiSeedRequest) String() string            { return proto.CompactTextString(m) }
func (*MultiSeedRequest) ProtoMessage()               {}
func (*MultiSeedRequest) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{1} }

func (m *MultiSeedRequest) GetUris() []string {
	if m != nil {
		return m.Uris
	}
	return nil
}

type Feedback struct {
	Uri              *string  `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	Type             *string  `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Timestamp        *float64 `protobuf:"fixed64,3,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Feedback) Reset()                    { *m = Feedback{} }
func (m *Feedback) String() string            { return proto.CompactTextString(m) }
func (*Feedback) ProtoMessage()               {}
func (*Feedback) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{2} }

func (m *Feedback) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *Feedback) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Feedback) GetTimestamp() float64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

type Tracks struct {
	Gids             []string    `protobuf:"bytes,1,rep,name=gids" json:"gids,omitempty"`
	Source           *string     `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	Identity         *string     `protobuf:"bytes,3,opt,name=identity" json:"identity,omitempty"`
	Tokens           []string    `protobuf:"bytes,4,rep,name=tokens" json:"tokens,omitempty"`
	Feedback         []*Feedback `protobuf:"bytes,5,rep,name=feedback" json:"feedback,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Tracks) Reset()                    { *m = Tracks{} }
func (m *Tracks) String() string            { return proto.CompactTextString(m) }
func (*Tracks) ProtoMessage()               {}
func (*Tracks) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{3} }

func (m *Tracks) GetGids() []string {
	if m != nil {
		return m.Gids
	}
	return nil
}

func (m *Tracks) GetSource() string {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return ""
}

func (m *Tracks) GetIdentity() string {
	if m != nil && m.Identity != nil {
		return *m.Identity
	}
	return ""
}

func (m *Tracks) GetTokens() []string {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func (m *Tracks) GetFeedback() []*Feedback {
	if m != nil {
		return m.Feedback
	}
	return nil
}

type Station struct {
	Id               *string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Title            *string  `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	TitleUri         *string  `protobuf:"bytes,3,opt,name=titleUri" json:"titleUri,omitempty"`
	Subtitle         *string  `protobuf:"bytes,4,opt,name=subtitle" json:"subtitle,omitempty"`
	SubtitleUri      *string  `protobuf:"bytes,5,opt,name=subtitleUri" json:"subtitleUri,omitempty"`
	ImageUri         *string  `protobuf:"bytes,6,opt,name=imageUri" json:"imageUri,omitempty"`
	LastListen       *float64 `protobuf:"fixed64,7,opt,name=lastListen" json:"lastListen,omitempty"`
	Seeds            []string `protobuf:"bytes,8,rep,name=seeds" json:"seeds,omitempty"`
	ThumbsUp         *int32   `protobuf:"varint,9,opt,name=thumbsUp" json:"thumbsUp,omitempty"`
	ThumbsDown       *int32   `protobuf:"varint,10,opt,name=thumbsDown" json:"thumbsDown,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Station) Reset()                    { *m = Station{} }
func (m *Station) String() string            { return proto.CompactTextString(m) }
func (*Station) ProtoMessage()               {}
func (*Station) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{4} }

func (m *Station) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Station) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *Station) GetTitleUri() string {
	if m != nil && m.TitleUri != nil {
		return *m.TitleUri
	}
	return ""
}

func (m *Station) GetSubtitle() string {
	if m != nil && m.Subtitle != nil {
		return *m.Subtitle
	}
	return ""
}

func (m *Station) GetSubtitleUri() string {
	if m != nil && m.SubtitleUri != nil {
		return *m.SubtitleUri
	}
	return ""
}

func (m *Station) GetImageUri() string {
	if m != nil && m.ImageUri != nil {
		return *m.ImageUri
	}
	return ""
}

func (m *Station) GetLastListen() float64 {
	if m != nil && m.LastListen != nil {
		return *m.LastListen
	}
	return 0
}

func (m *Station) GetSeeds() []string {
	if m != nil {
		return m.Seeds
	}
	return nil
}

func (m *Station) GetThumbsUp() int32 {
	if m != nil && m.ThumbsUp != nil {
		return *m.ThumbsUp
	}
	return 0
}

func (m *Station) GetThumbsDown() int32 {
	if m != nil && m.ThumbsDown != nil {
		return *m.ThumbsDown
	}
	return 0
}

type Rules struct {
	Js               *string `protobuf:"bytes,1,opt,name=js" json:"js,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Rules) Reset()                    { *m = Rules{} }
func (m *Rules) String() string            { return proto.CompactTextString(m) }
func (*Rules) ProtoMessage()               {}
func (*Rules) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{5} }

func (m *Rules) GetJs() string {
	if m != nil && m.Js != nil {
		return *m.Js
	}
	return ""
}

type StationResponse struct {
	Station          *Station    `protobuf:"bytes,1,opt,name=station" json:"station,omitempty"`
	Feedback         []*Feedback `protobuf:"bytes,2,rep,name=feedback" json:"feedback,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *StationResponse) Reset()                    { *m = StationResponse{} }
func (m *StationResponse) String() string            { return proto.CompactTextString(m) }
func (*StationResponse) ProtoMessage()               {}
func (*StationResponse) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{6} }

func (m *StationResponse) GetStation() *Station {
	if m != nil {
		return m.Station
	}
	return nil
}

func (m *StationResponse) GetFeedback() []*Feedback {
	if m != nil {
		return m.Feedback
	}
	return nil
}

type StationList struct {
	Stations         []*Station `protobuf:"bytes,1,rep,name=stations" json:"stations,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *StationList) Reset()                    { *m = StationList{} }
func (m *StationList) String() string            { return proto.CompactTextString(m) }
func (*StationList) ProtoMessage()               {}
func (*StationList) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{7} }

func (m *StationList) GetStations() []*Station {
	if m != nil {
		return m.Stations
	}
	return nil
}

type LikedPlaylist struct {
	Uri              *string `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LikedPlaylist) Reset()                    { *m = LikedPlaylist{} }
func (m *LikedPlaylist) String() string            { return proto.CompactTextString(m) }
func (*LikedPlaylist) ProtoMessage()               {}
func (*LikedPlaylist) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{8} }

func (m *LikedPlaylist) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func init() {
	proto.RegisterType((*RadioRequest)(nil), "Spotify.RadioRequest")
	proto.RegisterType((*MultiSeedRequest)(nil), "Spotify.MultiSeedRequest")
	proto.RegisterType((*Feedback)(nil), "Spotify.Feedback")
	proto.RegisterType((*Tracks)(nil), "Spotify.Tracks")
	proto.RegisterType((*Station)(nil), "Spotify.Station")
	proto.RegisterType((*Rules)(nil), "Spotify.Rules")
	proto.RegisterType((*StationResponse)(nil), "Spotify.StationResponse")
	proto.RegisterType((*StationList)(nil), "Spotify.StationList")
	proto.RegisterType((*LikedPlaylist)(nil), "Spotify.LikedPlaylist")
}

func init() { proto.RegisterFile("radio.proto", fileDescriptor15) }

var fileDescriptor15 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4f, 0x8f, 0xd3, 0x3e,
	0x10, 0x55, 0xd2, 0xa6, 0x4d, 0x26, 0xbf, 0x1f, 0x14, 0x0b, 0x81, 0x85, 0x10, 0x0a, 0x39, 0xa0,
	0x08, 0x41, 0x0f, 0x7b, 0xe5, 0x8a, 0x90, 0x90, 0x16, 0x84, 0x5c, 0xf6, 0x03, 0xa4, 0xb5, 0xb7,
	0xeb, 0x6d, 0x9a, 0x84, 0x8c, 0x23, 0xd4, 0xef, 0xc0, 0x7d, 0xbf, 0x2e, 0xf2, 0xd8, 0xce, 0x16,
	0x81, 0xe0, 0xf6, 0xde, 0x9b, 0x3f, 0x7e, 0x33, 0x63, 0xc8, 0x87, 0x5a, 0xea, 0x6e, 0xdd, 0x0f,
	0x9d, 0xe9, 0xd8, 0x72, 0xd3, 0x77, 0x46, 0x5f, 0x9f, 0xca, 0x1f, 0x11, 0xfc, 0x27, 0x6c, 0x40,
	0xa8, 0x6f, 0xa3, 0x42, 0xc3, 0x18, 0xcc, 0xc7, 0x41, 0x23, 0x8f, 0x8a, 0x59, 0x95, 0x09, 0xc2,
	0x56, 0xc3, 0xba, 0x31, 0x3c, 0x2e, 0xa2, 0x2a, 0x11, 0x84, 0xd9, 0x13, 0x58, 0x34, 0xaa, 0xdd,
	0x9b, 0x1b, 0x3e, 0x27, 0xd5, 0x33, 0xf6, 0x1c, 0x32, 0x34, 0xb5, 0xd1, 0x5d, 0xfb, 0x51, 0xf2,
	0xa4, 0x88, 0xaa, 0x4c, 0xdc, 0x0b, 0xec, 0x05, 0x40, 0x53, 0xa3, 0xf9, 0x3a, 0xd4, 0xbb, 0x03,
	0xf2, 0x05, 0xbd, 0x71, 0xa6, 0x94, 0xaf, 0x60, 0xf5, 0x69, 0x6c, 0x8c, 0xde, 0x28, 0x25, 0xff,
	0xe2, 0xa8, 0xfc, 0x0c, 0xe9, 0x07, 0xa5, 0xe4, 0xb6, 0xde, 0x1d, 0xd8, 0x0a, 0x66, 0xe3, 0xa0,
	0x79, 0x44, 0x6f, 0x59, 0x68, 0x2b, 0xcc, 0xa9, 0x57, 0xe4, 0x37, 0x13, 0x84, 0xad, 0x2f, 0xa3,
	0x8f, 0x0a, 0x4d, 0x7d, 0xec, 0xf9, 0xac, 0x88, 0xaa, 0x48, 0xdc, 0x0b, 0xe5, 0x5d, 0x04, 0x0b,
	0x67, 0xc1, 0x16, 0xef, 0xb5, 0x9c, 0x9e, 0xb3, 0xd8, 0x0e, 0x8b, 0xdd, 0x38, 0xec, 0x42, 0x4b,
	0xcf, 0xd8, 0x33, 0x48, 0xb5, 0x54, 0xad, 0xd1, 0xe6, 0x44, 0x3d, 0x33, 0x31, 0x71, 0x5b, 0x63,
	0xba, 0x83, 0x6a, 0x91, 0xcf, 0xa9, 0x93, 0x67, 0xec, 0x2d, 0xa4, 0xd7, 0xde, 0x3a, 0x4f, 0x8a,
	0x59, 0x95, 0x5f, 0x3c, 0x5a, 0xfb, 0x6b, 0xac, 0xc3, 0x4c, 0x62, 0x4a, 0x29, 0xef, 0x62, 0x58,
	0x6e, 0xdc, 0xfe, 0xd8, 0x03, 0x88, 0xb5, 0xf4, 0x83, 0xc6, 0x5a, 0xb2, 0xc7, 0x90, 0x18, 0x6d,
	0x9a, 0xe0, 0xca, 0x11, 0x6b, 0x8a, 0xc0, 0xd5, 0xa0, 0x83, 0xa9, 0xc0, 0x6d, 0x0c, 0xc7, 0xad,
	0x2b, 0x9a, 0xbb, 0x58, 0xe0, 0xac, 0x80, 0x3c, 0x60, 0x5b, 0xea, 0x6e, 0x77, 0x2e, 0xd1, 0xb8,
	0xc7, 0x7a, 0x4f, 0xe1, 0x85, 0x1f, 0xd7, 0xf3, 0x70, 0xd9, 0x4b, 0x8d, 0x46, 0xb5, 0x7c, 0x49,
	0x0b, 0x3e, 0x53, 0xac, 0x57, 0x54, 0x4a, 0x22, 0x4f, 0x69, 0x1b, 0x8e, 0x90, 0xd7, 0x9b, 0xf1,
	0xb8, 0xc5, 0xab, 0x9e, 0x67, 0xf4, 0x8f, 0x26, 0x6e, 0x3b, 0x3a, 0xfc, 0xbe, 0xfb, 0xde, 0x72,
	0xa0, 0xe8, 0x99, 0x52, 0x3e, 0x85, 0x44, 0x8c, 0x8d, 0x42, 0xbb, 0x96, 0x5b, 0x0c, 0x6b, 0xb9,
	0xc5, 0xb2, 0x81, 0x87, 0x7e, 0x63, 0x42, 0x61, 0xdf, 0xb5, 0xa8, 0xd8, 0x6b, 0x58, 0xfa, 0x4f,
	0x48, 0x79, 0xf9, 0xc5, 0x6a, 0xda, 0x79, 0x48, 0x0d, 0x09, 0xbf, 0x1c, 0x28, 0xfe, 0xf7, 0x81,
	0xde, 0x41, 0xee, 0x5b, 0xd8, 0x49, 0xd9, 0x1b, 0x48, 0x7d, 0x23, 0xf7, 0x85, 0xfe, 0xf4, 0xd4,
	0x94, 0x51, 0xbe, 0x84, 0xff, 0x2f, 0xf5, 0x41, 0xc9, 0x2f, 0x4d, 0x7d, 0x6a, 0x6c, 0xf9, 0x6f,
	0x9f, 0xf9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0xac, 0x82, 0x4e, 0xb8, 0x03, 0x00, 0x00,
}
