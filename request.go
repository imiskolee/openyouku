package openyouku

type Request struct {
	sysParam    *SysParams
	customParam interface{}
	Method      string
}
