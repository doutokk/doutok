namespace go auth

struct Token{
    1:string value,
    2:i64 expiration_time
}

service AuthService{
    Token GenerateToken(1:i64 user_id),
    bool ValidateToken(1:string token),
    bool RevokeToken(1:string token)

}