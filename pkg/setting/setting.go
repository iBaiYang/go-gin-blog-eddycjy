package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

//func NewSetting() (*Setting, error) {
//	vp := viper.New()
//	vp.SetConfigName("config")
//	vp.AddConfigPath("configs/")
//	vp.SetConfigType("yaml")
//	err := vp.ReadInConfig()
//	if err != nil {
//		return nil, err
//	}
//
//	return &Setting{vp}, nil
//}

/*配置读取*/
func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}

	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	//return &Setting{vp}, nil

	/*配置热更新*/
	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}

/*配置热更新*/
func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()
}
