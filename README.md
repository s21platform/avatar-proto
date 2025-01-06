# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [avatar.proto](#avatar-proto)
    - [Avatar](#avatar-Avatar)
    - [DeleteSocietyAvatarIn](#avatar-DeleteSocietyAvatarIn)
    - [DeleteUserAvatarIn](#avatar-DeleteUserAvatarIn)
    - [GetAllSocietyAvatarsIn](#avatar-GetAllSocietyAvatarsIn)
    - [GetAllSocietyAvatarsOut](#avatar-GetAllSocietyAvatarsOut)
    - [GetAllUserAvatarsIn](#avatar-GetAllUserAvatarsIn)
    - [GetAllUserAvatarsOut](#avatar-GetAllUserAvatarsOut)
    - [SetSocietyAvatarIn](#avatar-SetSocietyAvatarIn)
    - [SetSocietyAvatarOut](#avatar-SetSocietyAvatarOut)
    - [SetUserAvatarIn](#avatar-SetUserAvatarIn)
    - [SetUserAvatarOut](#avatar-SetUserAvatarOut)
  
    - [AvatarService](#avatar-AvatarService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="avatar-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## avatar.proto



<a name="avatar-Avatar"></a>

### Avatar



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| link | [string](#string) |  |  |






<a name="avatar-DeleteSocietyAvatarIn"></a>

### DeleteSocietyAvatarIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| avatar_id | [int32](#int32) |  |  |






<a name="avatar-DeleteUserAvatarIn"></a>

### DeleteUserAvatarIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| avatar_id | [int32](#int32) |  |  |






<a name="avatar-GetAllSocietyAvatarsIn"></a>

### GetAllSocietyAvatarsIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |






<a name="avatar-GetAllSocietyAvatarsOut"></a>

### GetAllSocietyAvatarsOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| avatar_list | [Avatar](#avatar-Avatar) | repeated |  |






<a name="avatar-GetAllUserAvatarsIn"></a>

### GetAllUserAvatarsIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |






<a name="avatar-GetAllUserAvatarsOut"></a>

### GetAllUserAvatarsOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| avatar_list | [Avatar](#avatar-Avatar) | repeated |  |






<a name="avatar-SetSocietyAvatarIn"></a>

### SetSocietyAvatarIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| filename | [string](#string) |  |  |
| batch | [bytes](#bytes) |  |  |






<a name="avatar-SetSocietyAvatarOut"></a>

### SetSocietyAvatarOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| link | [string](#string) |  |  |






<a name="avatar-SetUserAvatarIn"></a>

### SetUserAvatarIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| filename | [string](#string) |  |  |
| batch | [bytes](#bytes) |  |  |






<a name="avatar-SetUserAvatarOut"></a>

### SetUserAvatarOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| link | [string](#string) |  |  |





 

 

 


<a name="avatar-AvatarService"></a>

### AvatarService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SetUserAvatar | [SetUserAvatarIn](#avatar-SetUserAvatarIn) stream | [SetUserAvatarOut](#avatar-SetUserAvatarOut) |  |
| GetAllUserAvatars | [GetAllUserAvatarsIn](#avatar-GetAllUserAvatarsIn) | [GetAllUserAvatarsOut](#avatar-GetAllUserAvatarsOut) |  |
| DeleteUserAvatar | [DeleteUserAvatarIn](#avatar-DeleteUserAvatarIn) | [Avatar](#avatar-Avatar) |  |
| SetSocietyAvatar | [SetSocietyAvatarIn](#avatar-SetSocietyAvatarIn) stream | [SetSocietyAvatarOut](#avatar-SetSocietyAvatarOut) |  |
| GetAllSocietyAvatars | [GetAllSocietyAvatarsIn](#avatar-GetAllSocietyAvatarsIn) | [GetAllSocietyAvatarsOut](#avatar-GetAllSocietyAvatarsOut) |  |
| DeleteSocietyAvatar | [DeleteSocietyAvatarIn](#avatar-DeleteSocietyAvatarIn) | [Avatar](#avatar-Avatar) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

