---
title: 默认模块
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.30"

---

# 默认模块

Base URLs:

# Authentication

- HTTP Authentication, scheme: bearer

# API 文档

## POST 对话补全

POST /chat/completions

根据输入的上下文，来让模型补全对话内容。

> Body 请求参数

```json
{
  "messages": [
    {
      "content": "You are a helpful assistant",
      "role": "system"
    },
    {
      "content": "你是谁？",
      "role": "user"
    }
  ],
  "model": "deepseek-chat",
  "stream": true
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» messages|body|[oneOf]| 是 ||对话的消息列表|
|»» *anonymous*|body|object| 否 | System message|none|
|»»» content|body|string| 是 ||system 消息的内容|
|»»» role|body|string| 是 ||该消息的发起角色，其值为 `system`|
|»»» name|body|string| 否 ||可以选填的参与者的名称，为模型提供信息以区分相同角色的参与者|
|»» *anonymous*|body|object| 否 | User message|none|
|»»» content|body|string| 是 ||user 消息的内容|
|»»» role|body|string| 是 ||该消息的发起角色，其值为 `user`|
|»»» name|body|string| 否 ||可以选填的参与者的名称，为模型提供信息以区分相同角色的参与者|
|»» *anonymous*|body|object| 否 | Assistant message|none|
|»»» content|body|string| 是 ||assistant 消息的内容|
|»»» role|body|string| 是 ||该消息的发起角色，其值为 `assistant`|
|»»» name|body|string| 否 ||可以选填的参与者的名称，为模型提供信息以区分相同角色的参与者。|
|»» *anonymous*|body|object| 否 | Tool message|none|
|»»» content|body|string| 是 ||tool 消息的内容|
|»»» role|body|string| 是 ||该消息的发起角色，其值为 `tool`|
|»»» tool_call_id|body|string| 是 ||此消息所响应的 tool call 的 ID|
|» model|body|string| 是 ||使用的 AI 模型|
|» frequency_penalty|body|number¦null| 否 ||none|
|» max_tokens|body|integer¦null| 否 ||none|
|» presence_penalty|body|number¦null| 否 ||none|
|» response_format|body|object¦null| 否 ||none|
|»» type|body|string¦null| 否 ||none|
|» stop|body|any| 否 ||none|
|»» *anonymous*|body|string¦null| 否 ||none|
|»» *anonymous*|body|[string]¦null| 否 ||none|
|» stream|body|boolean¦null| 否 ||是否启用流式传输|
|» stream_options|body|object¦null| 否 ||none|
|»» include_usage|body|boolean| 否 ||none|
|» temperature|body|number¦null| 否 ||none|
|» top_p|body|number¦null| 否 ||none|
|» tools|body|null| 否 ||none|
|» tool_choice|body|any| 否 ||none|
|»» *anonymous*|body|string| 否 ||none|
|»» *anonymous*|body|object| 否 ||none|
|» logprobs|body|boolean¦null| 否 ||none|
|» top_logprobs|body|integer¦null| 否 ||none|

#### 枚举值

|属性|值|
|---|---|
|» model|deepseek-chat|
|» model|deepseek-reasoner|

> 返回示例

> 200 Response

```json
{
  "id": "e137bb42-7580-4cb8-88ba-825209cf966b",
  "choices": [
    {
      "index": 0,
      "message": {
        "role": "assistant",
        "content": "Hello! How can I assist you today? 😊"
      },
      "logprobs": null,
      "finish_reason": "stop"
    }
  ],
  "created": 1739112811,
  "model": "deepseek-chat",
  "system_fingerprint": "fp_3a5790e1b4",
  "object": "chat.completion",
  "usage": {
    "prompt_tokens": 9,
    "completion_tokens": 11,
    "total_tokens": 20,
    "prompt_tokens_details": {
      "cached_tokens": 0
    },
    "prompt_cache_hit_tokens": 0,
    "prompt_cache_miss_tokens": 9
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK, 返回一个 `chat completion` 对象。|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» id|string|true|none||none|
|» choices|[object]|true|none||none|
|»» finish_reason|string|false|none||none|
|»» index|integer|false|none||none|
|»» message|object|false|none||none|
|»»» content|string¦null|true|none||none|
|»»» reasoning_content|string¦null|false|none||none|
|»»» tool_calls|[object]|false|none||none|
|»»»» id|string|false|none||none|
|»»»» type|string|false|none||none|
|»»»» function|object|false|none||none|
|»»»»» name|string|true|none||none|
|»»»»» arguments|string|true|none||none|
|»»» role|string|true|none||none|
|»» logprobs|object¦null|true|none||none|
|»»» content|[object]¦null|true|none||none|
|»»»» token|string|true|none||none|
|»»»» logprob|number|true|none||none|
|»»»» bytes|[integer]¦null|true|none||none|
|»»»» top_logprobs|[object]|true|none||none|
|»»»»» token|string|false|none||none|
|»»»»» logprob|integer|false|none||none|
|»»»»» bytes|[integer]|false|none||none|
|» created|integer|true|none||none|
|» model|string|true|none||none|
|» system_fingerprint|string|true|none||none|
|» object|string|true|none||none|
|» usage|object|false|none||none|
|»» completion_tokens|integer|true|none||none|
|»» prompt_tokens|integer|true|none||none|
|»» prompt_cache_hit_tokens|integer|true|none||none|
|»» prompt_cache_miss_tokens|integer|true|none||none|
|»» total_tokens|integer|true|none||none|
|»» completion_tokens_details|object|false|none||none|
|»»» reasoning_tokens|integer|true|none||none|

## POST FIM 补全（Beta）

POST /beta/completions

FIM（Fill-In-the-Middle）补全 API。

用户需要设置 base_url="https://api.deepseek.com/beta" 来使用此功能。

> Body 请求参数

```json
{
  "model": "deepseek-chat",
  "prompt": "Once upon a time, ",
  "echo": false,
  "frequency_penalty": 0,
  "logprobs": 0,
  "max_tokens": 1024,
  "presence_penalty": 0,
  "stop": null,
  "stream": false,
  "stream_options": null,
  "suffix": null,
  "temperature": 1,
  "top_p": 1
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» model|body|string| 是 ||none|
|» prompt|body|string| 是 ||none|
|» echo|body|boolean¦null| 否 ||none|
|» frequency_penalty|body|number¦null| 否 ||none|
|» logprobs|body|integer¦null| 否 ||none|
|» max_tokens|body|integer¦null| 否 ||none|
|» presence_penalty|body|number¦null| 否 ||none|
|» stop|body|null| 否 ||none|
|» stream|body|boolean¦null| 否 ||none|
|» stream_options|body|null| 否 ||none|
|» suffix|body|string¦null| 否 ||none|
|» temperature|body|number¦null| 否 ||none|
|» top_p|body|number¦null| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "id": "string",
  "choices": [
    {
      "finish_reason": "stop",
      "index": 0,
      "logprobs": {
        "text_offset": [
          0
        ],
        "token_logprobs": [
          0
        ],
        "tokens": [
          "string"
        ],
        "top_logprobs": [
          {}
        ]
      },
      "text": "string"
    }
  ],
  "created": 0,
  "model": "string",
  "system_fingerprint": "string",
  "object": "text_completion",
  "usage": {
    "completion_tokens": 0,
    "prompt_tokens": 0,
    "prompt_cache_hit_tokens": 0,
    "prompt_cache_miss_tokens": 0,
    "total_tokens": 0,
    "completion_tokens_details": {
      "reasoning_tokens": 0
    }
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» id|string|true|none||none|
|» choices|[object]|true|none||none|
|»» finish_reason|string|false|none||none|
|»» index|integer|false|none||none|
|»» logprobs|object¦null|false|none||none|
|»»» text_offset|[integer]¦null|false|none||none|
|»»» token_logprobs|[number]¦null|false|none||none|
|»»» tokens|[string]¦null|false|none||none|
|»»» top_logprobs|[object]¦null|false|none||none|
|»» text|string|false|none||none|
|» created|integer|true|none||none|
|» model|string|true|none||none|
|» system_fingerprint|string|false|none||none|
|» object|string|true|none||none|
|» usage|object|false|none||none|
|»» completion_tokens|integer|true|none||none|
|»» prompt_tokens|integer|true|none||none|
|»» prompt_cache_hit_tokens|integer|true|none||none|
|»» prompt_cache_miss_tokens|integer|true|none||none|
|»» total_tokens|integer|true|none||none|
|»» completion_tokens_details|object|false|none||none|
|»»» reasoning_tokens|integer|false|none||none|

## GET 列出模型

GET /models

列出可用的模型列表，并提供相关模型的基本信息。请前往[模型 & 价格](https://api-docs.deepseek.com/zh-cn/quick_start/pricing)查看当前支持的模型列表

> 返回示例

> 200 Response

```json
{
  "object": "list",
  "data": [
    {
      "id": "deepseek-chat",
      "object": "model",
      "owned_by": "deepseek"
    },
    {
      "id": "deepseek-reasoner",
      "object": "model",
      "owned_by": "deepseek"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK, 返回模型列表|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» object|string|true|none||none|
|» data|[object]|true|none||none|
|»» id|string|true|none||none|
|»» object|string|true|none||none|
|»» owned_by|string|true|none||none|

## GET 查询余额

GET /user/balance

> 返回示例

> 200 Response

```json
{
  "is_available": true,
  "balance_infos": [
    {
      "currency": "CNY",
      "total_balance": "110.00",
      "granted_balance": "10.00",
      "topped_up_balance": "100.00"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» is_available|boolean|true|none||none|
|» balance_infos|[object]|true|none||none|
|»» currency|string|false|none||none|
|»» total_balance|string|false|none||none|
|»» granted_balance|string|false|none||none|
|»» topped_up_balance|string|false|none||none|

# 数据模型

