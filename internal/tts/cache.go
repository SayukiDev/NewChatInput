package tts

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/klauspost/compress/zstd"
)

func (t *TTS) getCachePath(text string) string {
	name := fmt.Sprintf("%s:%d", text, t.o.NowSpacker)
	sum := sha256.Sum256([]byte(name))
	hash := hex.EncodeToString(sum[:16])
	return path.Join(t.o.CachePath, hash+".wav")
}

func (t *TTS) writeCache(text string, b []byte) (retErr error) {
	cachePath := t.getCachePath(text)
	openFile, err := os.OpenFile(cachePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func() {
		openFile.Close()
		if retErr != nil {
			os.Remove(cachePath)
		}
	}()
	w, err := zstd.NewWriter(openFile)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	if err != nil {
		w.Close()
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}
	return nil
}

func (t *TTS) readCache(text string) ([]byte, error) {
	getCachePath := t.getCachePath(text)
	openFile, err := os.Open(getCachePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	defer openFile.Close()
	r, err := zstd.NewReader(openFile)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return b, nil
}
