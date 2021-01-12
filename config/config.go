package config

const baseAction = "a:"
const baseUser = "u:"
const basePolicy = "p:"
const baseResource = "r:"
const baseSubject = "s:"

const EffectAllow = "allow"

// Users
const UserAdmin = baseUser + "0"
const User1 = baseUser + "1"
const User2 = baseUser + "2"
const User3 = baseUser + "3"
const User4 = baseUser + "4"
const User5 = baseUser + "5"

// Product A
const ProductAActionRead = baseAction + "read"
const ProductAActionEdit = baseAction + "edit"
const ProductAActionDelete = baseAction + "delete"

const ProductARoleReader = "reader:"
const ProductARoleEditor = "editor:"
const ProductARoleManager = "manager:"

const ProductASubjectAdmin = baseSubject + "admin"

// Generated
const ProductAProject1 = "project:1951823"
const ProductAProject1Resource = baseResource + ProductAProject1
const ProductAProject1PolicyManage = basePolicy + ProductARoleManager + ProductAProject1
const ProductAProject1PolicyEdit = basePolicy + ProductARoleEditor + ProductAProject1
const ProductAProject1PolicyRead = basePolicy + ProductARoleReader + ProductAProject1
const ProductAProject1SubjectManage = baseSubject + ProductARoleManager + ProductAProject1
const ProductAProject1SubjectEdit = baseSubject + ProductARoleEditor + ProductAProject1
const ProductAProject1SubjectRead = baseSubject + ProductARoleReader + ProductAProject1

const ProductAProject2 = "project:5725100"
const ProductAProject2Resource = baseResource + ProductAProject2
const ProductAProject2PolicyManage = basePolicy + ProductARoleManager + ProductAProject2
const ProductAProject2PolicyEdit = basePolicy + ProductARoleEditor + ProductAProject2
const ProductAProject2PolicyRead = basePolicy + ProductARoleReader + ProductAProject2
const ProductAProject2SubjectManage = baseSubject + ProductARoleManager + ProductAProject2
const ProductAProject2SubjectEdit = baseSubject + ProductARoleEditor + ProductAProject2
const ProductAProject2SubjectRead = baseSubject + ProductARoleReader + ProductAProject2
