import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// リクエストデータを格納するための構造体
type CreateRequestParam struct {
	Task string `json:"task" binding:"required,max=60"`
}

func (t *todoHandler) Create(c *gin.Context) {
	var req CreateRequestParam
	// リクエストパラメータを構造体（CreateRequestParam）にマッピング
	if err := c.ShouldBindJSON(&req); err != nil {
		// バリデーションエラーがあった場合はエラーを返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// usecase の呼び出し
	err := t.usecase.Create(req.Task)
	if err != nil {
		// エラーがあった場合はエラーレスポンスを返却
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	// レスポンスを返却
	c.JSON(http.StatusCreated, nil)
}

// パス(todo/:id)に指定されたパラメータを格納するための構造体
type UpdateRequestPathParam struct {
	ID int `uri:"id"`
}

// リクエストデータ(body)を格納するための構造体
type UpdateRequestBodyParam struct {
	Task   string           `json:"task" binding:"required,max=60"`
	Status model.TaskStatus `json:"status" binding:"required,task_status"`
}

func (t *todoHandler) Update(c *gin.Context) {
	var pathParam UpdateRequestPathParam
	var bodyParam UpdateRequestBodyParam
	// パスパラメータを構造体（UpdateRequestPathParam）にマッピング
	if err := c.ShouldBindUri(&pathParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// リクエストパラメータ(body)を構造体（UpdateRequestBodyParam）にマッピング
	if err := c.ShouldBindJSON(&bodyParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// usecase の呼び出し
	if err := t.usecase.Update(pathParam.ID, bodyParam.Task, bodyParam.Status); err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	// レスポンスを返却
	c.JSON(http.StatusNoContent, nil)
}

// パス(todo/:id)に指定されたパラメータを格納するための構造体
type FindRequestParam struct {
	ID int `uri:"id" binding:"required"`
}

func (t *todoHandler) Find(c *gin.Context) {
	var req FindRequestParam
	// パスパラメータを構造体（FindRequestParam）にマッピング

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// usecase の呼び出し
	res, err := t.usecase.Find(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		// 検索結果(res = nil）が存在しなかった場合は 404 not found を返す
		c.JSON(http.StatusNotFound, nil)
		return
	}
	// レスポンスを返却
	c.JSON(http.StatusOK, res)
}

func (t *todoHandler) FindAll(c *gin.Context) {
	res, err := t.usecase.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// パス(todo/:id)に指定されたパラメータを格納するための構造体
type DeleteRequestParam struct {
	ID int `uri:"id"`
}

func (t *todoHandler) Delete(c *gin.Context) {
	var req DeleteRequestParam
	// パスパラメータを構造体（DeleteRequestParam）にマッピング
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// usecase の呼び出し
	if err := t.usecase.Delete(req.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// レスポンスを返却
	c.JSON(http.StatusNoContent, nil)
}
