package gfs

import (
	"time"
)

//------ ChunkServer

type PushDataAndForwardArg struct {
    Handle    ChunkHandle
	Data      []byte
	ForwardTo []ServerAddress
}

type PushDataAndForwardReply struct {
	DataID DataBufferID
}

type ForwardDataArg struct {
	DataID DataBufferID
	Data   []byte
}

type ForwardDataReply struct {}

type CreateChunkArg struct {
	Handle ChunkHandle
}
type CreateChunkReply struct {}

type WriteChunkArg struct {
	DataID      DataBufferID
	Offset      Offset
	Secondaries []ServerAddress
}
type WriteChunkReply struct {}

type AppendChunkArg struct {
	DataID      DataBufferID
	Secondaries []ServerAddress
}
type AppendChunkReply struct {
    Offset Offset
}

type ApplyMutationArg struct {
    Mtype   MutationType
    Version ChunkVersion
	DataID  DataBufferID
    Offset  Offset
}
type ApplyMutationReply struct {}

type PadChunkArg struct {
	Handle ChunkHandle
}
type PadChunkReply struct {}


//------ Master

type HeartbeatArg struct {
	Address         ServerAddress // chunkserver address
	LeaseExtensions []ChunkHandle // leases to be extended
}

type HeartbeatReply struct{}

type GetPrimaryAndSecondariesArg struct {
	Handle ChunkHandle
}

type GetPrimaryAndSecondariesReply struct {
	Primary     ServerAddress
	Expire      time.Time
	Secondaries []ServerAddress
}

type ExtendLeaseArg struct {
	Handle  ChunkHandle
	Address ServerAddress
}

type ExtendLeaseReply struct {
	Expire time.Time
}

type GetReplicasArg struct {
	Handle ChunkHandle
}

type GetReplicasReply struct {
	Locations []ServerAddress
}

type GetFileInfoArg struct {
    Path Path
}

type GetFileInfoReply struct {
    IsDir  bool
    Length int64
    Chunks int64
}

type CreateFileArg struct {
    Path Path
}

type CreateFileReply struct {}


type GetChunkHandleArg struct {
	Path  Path
	Index ChunkIndex
}

type GetChunkHandleReply struct {
	Handle ChunkHandle
}