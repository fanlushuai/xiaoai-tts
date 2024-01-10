# xiaoai-tts
Xiaoai speaker customizes the text to read aloud.


## 安装

```bash
go get xxxxxxx
```

## Example

```golang
m := &models.MiAccount{
		User: "xxxx",
		Pwd:  "xxxx",
	}
//new a xioaxi client
xiaoai := service.NewXiaoAi(m)
//
xiaoai.Say("hello")
//
xiaoai.GetDevice() []models.DeviceInfo
//
xiaoai.UseDevice(index int16)
//
xiaoai.Say(text string)
//
xiaoai.SetVolume(volume int8)
//
xiaoai.GetVolume() string
//
xiaoai.Play()
//
xiaoai.Pause()
//
xiaoai.Prev()
//
xiaoai.Next()
//
xiaoai.TogglePlayState()
//
xiaoai.GetStatus() *models.Info
//
xiaoai.PlayUrl(url string)
```

## Docker
1. Add .env on the root dir

```bash
echo "User=xx\nPwd=xxx"> .env
```
2. Up it
```bash
docker-compose -f docker-compose.yml up
```
3. Test it
```bash
curl --location 'http://127.0.0.1:8848/xiaoai/speak' \
--form 'msg="a"'
