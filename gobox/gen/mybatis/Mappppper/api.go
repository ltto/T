package Mappppper

import (
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/ltto/T/echoT"
	"github.com/ltto/T/echoT/vo"
)

func init() {
	//******************Albums******************//
	//Albums get one
	echoT.R(echoT.RouterInfo{Mapping: "/albums", HttpMethod: http.MethodGet,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf(Albums{})},
		Do: func(c echo.Context, res struct {
			ID int `query:"Cid"`
		}) vo.Result {
			if one, err := AlbumsMP.SelectByID(res.ID); err != nil {
				return *vo.Err(err)
			} else {
				return *vo.Success(one)
			}
		},
	})
	//Albums get list
	echoT.R(echoT.RouterInfo{Mapping: "/albumss", HttpMethod: http.MethodGet,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf([]Albums{})},
		Do: func(c echo.Context, res *vo.Page) vo.Result {
			list, err := AlbumsMP.SelectLimit(res.Limit())
			if err != nil {
				return *vo.Err(err)
			}
			count, err := AlbumsMP.SelectCount()
			if err != nil {
				return *vo.Err(err)
			}
			res.Count = count
			return *vo.List(list, res)
		},
	})
	//Albums save
	echoT.R(echoT.RouterInfo{Mapping: "/albums/save", HttpMethod: http.MethodPost,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf(Albums{})},
		Do: func(c echo.Context, res Albums) vo.Result {
			if err := AlbumsMP.Save(res); err != nil {
				return *vo.Err(err)
			}
			return *vo.Success(res)
		},
	})
	//Albums update
	echoT.R(echoT.RouterInfo{Mapping: "/albums", HttpMethod: http.MethodPost,
		Auth:         false,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf(Albums{})},
		Do: func(c echo.Context, res Albums) vo.Result {
			if err := AlbumsMP.UpdateByID(res); err != nil {
				return *vo.Err(err)
			}
			return *vo.Success(res)
		},
	})
	//------------------------

	//******************BizProperty******************//
	//BizProperty get one
	echoT.R(echoT.RouterInfo{Mapping: "/bizproperty", HttpMethod: http.MethodGet,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf(BizProperty{})},
		Do: func(c echo.Context, res struct {
			ID int `query:"Id"`
		}) vo.Result {
			if one, err := BizPropertyMP.SelectByID(res.ID); err != nil {
				return *vo.Err(err)
			} else {
				return *vo.Success(one)
			}
		},
	})
	//BizProperty get list
	echoT.R(echoT.RouterInfo{Mapping: "/bizpropertys", HttpMethod: http.MethodGet,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf([]BizProperty{})},
		Do: func(c echo.Context, res *vo.Page) vo.Result {
			list, err := BizPropertyMP.SelectLimit(res.Limit())
			if err != nil {
				return *vo.Err(err)
			}
			count, err := BizPropertyMP.SelectCount()
			if err != nil {
				return *vo.Err(err)
			}
			res.Count = count
			return *vo.List(list, res)
		},
	})
	//BizProperty save
	echoT.R(echoT.RouterInfo{Mapping: "/bizproperty/save", HttpMethod: http.MethodPost,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf(BizProperty{})},
		Do: func(c echo.Context, res BizProperty) vo.Result {
			if err := BizPropertyMP.Save(res); err != nil {
				return *vo.Err(err)
			}
			return *vo.Success(res)
		},
	})
	//BizProperty update
	echoT.R(echoT.RouterInfo{Mapping: "/bizproperty", HttpMethod: http.MethodPost,
		Auth:         false,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf(BizProperty{})},
		Do: func(c echo.Context, res BizProperty) vo.Result {
			if err := BizPropertyMP.UpdateByID(res); err != nil {
				return *vo.Err(err)
			}
			return *vo.Success(res)
		},
	})
	//------------------------

	//******************Photos******************//
	//Photos get one
	echoT.R(echoT.RouterInfo{Mapping: "/photos", HttpMethod: http.MethodGet,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf(Photos{})},
		Do: func(c echo.Context, res struct {
			ID int `query:"Id"`
		}) vo.Result {
			if one, err := PhotosMP.SelectByID(res.ID); err != nil {
				return *vo.Err(err)
			} else {
				return *vo.Success(one)
			}
		},
	})
	//Photos get list
	echoT.R(echoT.RouterInfo{Mapping: "/photoss", HttpMethod: http.MethodGet,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf([]Photos{})},
		Do: func(c echo.Context, res *vo.Page) vo.Result {
			list, err := PhotosMP.SelectLimit(res.Limit())
			if err != nil {
				return *vo.Err(err)
			}
			count, err := PhotosMP.SelectCount()
			if err != nil {
				return *vo.Err(err)
			}
			res.Count = count
			return *vo.List(list, res)
		},
	})
	//Photos save
	echoT.R(echoT.RouterInfo{Mapping: "/photos/save", HttpMethod: http.MethodPost,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf(Photos{})},
		Do: func(c echo.Context, res Photos) vo.Result {
			if err := PhotosMP.Save(res); err != nil {
				return *vo.Err(err)
			}
			return *vo.Success(res)
		},
	})
	//Photos update
	echoT.R(echoT.RouterInfo{Mapping: "/photos", HttpMethod: http.MethodPost,
		Auth:         false,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf(Photos{})},
		Do: func(c echo.Context, res Photos) vo.Result {
			if err := PhotosMP.UpdateByID(res); err != nil {
				return *vo.Err(err)
			}
			return *vo.Success(res)
		},
	})
	//------------------------

	//******************BizActivity******************//
	//BizActivity get one
	echoT.R(echoT.RouterInfo{Mapping: "/bizactivity", HttpMethod: http.MethodGet,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf(BizActivity{})},
		Do: func(c echo.Context, res struct {
			ID int `query:"Id"`
		}) vo.Result {
			if one, err := BizActivityMP.SelectByID(res.ID); err != nil {
				return *vo.Err(err)
			} else {
				return *vo.Success(one)
			}
		},
	})
	//BizActivity get list
	echoT.R(echoT.RouterInfo{Mapping: "/bizactivitys", HttpMethod: http.MethodGet,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf([]BizActivity{})},
		Do: func(c echo.Context, res *vo.Page) vo.Result {
			list, err := BizActivityMP.SelectLimit(res.Limit())
			if err != nil {
				return *vo.Err(err)
			}
			count, err := BizActivityMP.SelectCount()
			if err != nil {
				return *vo.Err(err)
			}
			res.Count = count
			return *vo.List(list, res)
		},
	})
	//BizActivity save
	echoT.R(echoT.RouterInfo{Mapping: "/bizactivity/save", HttpMethod: http.MethodPost,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf(BizActivity{})},
		Do: func(c echo.Context, res BizActivity) vo.Result {
			if err := BizActivityMP.Save(res); err != nil {
				return *vo.Err(err)
			}
			return *vo.Success(res)
		},
	})
	//BizActivity update
	echoT.R(echoT.RouterInfo{Mapping: "/bizactivity", HttpMethod: http.MethodPost,
		Auth:         false,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf(BizActivity{})},
		Do: func(c echo.Context, res BizActivity) vo.Result {
			if err := BizActivityMP.UpdateByID(res); err != nil {
				return *vo.Err(err)
			}
			return *vo.Success(res)
		},
	})
	//------------------------

	//******************Users******************//
	//Users get one
	echoT.R(echoT.RouterInfo{Mapping: "/users", HttpMethod: http.MethodGet,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf(Users{})},
		Do: func(c echo.Context, res struct {
			ID int `query:"Id"`
		}) vo.Result {
			if one, err := UsersMP.SelectByID(res.ID); err != nil {
				return *vo.Err(err)
			} else {
				return *vo.Success(one)
			}
		},
	})
	//Users get list
	echoT.R(echoT.RouterInfo{Mapping: "/userss", HttpMethod: http.MethodGet,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf([]Users{})},
		Do: func(c echo.Context, res *vo.Page) vo.Result {
			list, err := UsersMP.SelectLimit(res.Limit())
			if err != nil {
				return *vo.Err(err)
			}
			count, err := UsersMP.SelectCount()
			if err != nil {
				return *vo.Err(err)
			}
			res.Count = count
			return *vo.List(list, res)
		},
	})
	//Users save
	echoT.R(echoT.RouterInfo{Mapping: "/users/save", HttpMethod: http.MethodPost,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf(Users{})},
		Do: func(c echo.Context, res Users) vo.Result {
			if err := UsersMP.Save(res); err != nil {
				return *vo.Err(err)
			}
			return *vo.Success(res)
		},
	})
	//Users update
	echoT.R(echoT.RouterInfo{Mapping: "/users", HttpMethod: http.MethodPost,
		Auth:         false,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf(Users{})},
		Do: func(c echo.Context, res Users) vo.Result {
			if err := UsersMP.UpdateByID(res); err != nil {
				return *vo.Err(err)
			}
			return *vo.Success(res)
		},
	})
	//------------------------

}
