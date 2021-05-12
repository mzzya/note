# umi使用

## 分环境配置

```json
//package.json
{
    "scripts": {
        "start": "cross-env UMI_ENV=dev umi dev",
        "start:test": "cross-env UMI_ENV=test umi dev",
        "start:uat": "cross-env UMI_ENV=uat umi dev",
        "start:prd": "cross-env UMI_ENV=prd umi dev",
        "build": "cross-env UMI_ENV=dev umi build",
        "build:test": "cross-env UMI_ENV=test umi build",
        "build:uat": "cross-env UMI_ENV=uat umi build",
        "build:prd": "cross-env UMI_ENV=prd umi build",
    }
}
```

```ts
//.umirc.test.ts
//.umirc.local.ts中不要写任何东西，它的优先级最高，否则会冲掉环境标识的配置
import { defineConfig } from 'umi'

export default defineConfig({
  proxy: {
    '/api': {
      'target': 'https://sso.example.com',
      'changeOrigin': true,
      // 'pathRewrite': { '^/api': '' },
    },
  },
  define: {
    'process.env.ssoOrigin': "https://sso.example.com",
  }
})
```
