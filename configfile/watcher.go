package configfile

type FileWatcher func(callback func()) (unwatcher func())
