# FIBER TEMPLATE'S NOTE

Note for Fiber Template

## GORM NOTES

- If you want to use field alias (with operator `AS`)

  - Step 1: Declare in Model Struct's Field

    ```go
    type Something struct {
      // Migrated Fields
      ID        uint   `gorm:"primaryKey"`
      XID       string `gorm:"column:xid"`
      ...
      CreatedAt time.Time
      UpdatedAt time.Time
      DeletedAt gorm.DeletedAt `gorm:"index"`

      // Relations
      ...

      // Not Migrated Fields
      SomethingXID string `gorm:"column:something_xid;<-:false;-:migration"`
      ...
    }
    ```

  - Step 2: Type Select Statement Like This

    ```go
    ...
    &types.QuerySQL{
      Selects: []types.SelectQuerySQLOperation{
        {Field: "id"},
        {Field: "xid", Alias: "something_xid"},
        ...
        {Field: "created_at"},
        {Field: "updated_at"},
      },
    }
    ...
    ```

- If you want to do raw query

  - Step 1: Declare New Struct Payload and Result Next to the Model

    ```go
    type Something struct {
      // Migrated Fields
      ID        uint   `gorm:"primaryKey"`
      XID       string `gorm:"column:xid"`
      ...
      CreatedAt time.Time
      UpdatedAt time.Time
      DeletedAt gorm.DeletedAt `gorm:"index"`

      // Relations
      ...
    }

    type SomethingDoPayload struct {
      Field string
    }

    type SomethingDoResult struct {
      Field string `gorm:"column:field"`
    }
    ```

  - Step 2: Add Method name to the struct interface

    ```go
    type ISomethingDatabaseSQLRepository interface {
      ...

      DoSomething(payload *sql.SomethingDoPayload) (*sql.SomethingDoResult, error)
    }

    type SomethingDatabaseSQLRepository struct {
      db        *gorm.DB
      model     sql.Something
      dialect   pkg_constant.DialectDatabaseSQL
    }

    func InitMysqlSomethingDatabaseSQLRepository(db config.MysqlDatabaseSQLConnection) ISomethingDatabaseSQLRepository {
      return &SomethingDatabaseSQLRepository{
        db:        db,
        model:     sql.Something{},
        dialect:   pkg_constant.MYSQL,
      }
    }
    ```

  - Step 3: Write the code implementation in the struct method

  ```go
  func (r *SomethingDatabaseSQLRepository) DoSomething(payload *sql.SomethingDoPayload) (*sql.SomethingDoResult, error) {
    result := new(sql.SomethingDoResult)

    q := r.db

    if err := q.Raw(
      `
        SELECT
          s.xid
        FROM
          somethings s
        WHERE
          s.xid = ?
      `,
      payload.XID,
    ).Scan(result).Error; err != nil {
      return nil, err
    }

    return result, nil
  }
  ```
