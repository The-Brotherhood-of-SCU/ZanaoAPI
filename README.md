# ZanaoAPI

赞哦校园集市API代理

## 接口

### 1.主页帖子

`/home?school=scu&from=0`

获取主页的10条帖子，from是时间戳，分页依据

返回示例：

```json
{
  "data": {
    "cert_show_open": false,
    "list": [
      {
        "c_count": 1,
        "cate_id": "101",
        "cate_name": "二手闲置",
        "cert_show": "0",
        "color_show_info": {
          "bottom": "#0038FF",
          "top": "#8600FF"
        },
        "comment_list": [
          {
            "comment_id": "2048974424",
            "content": "cy",
            "headimgurl": "//b1.cdn.zanao.com/upload/2025/04/06/qbqtfhcsd69jxq8.jpeg@!sm_w100_h100",
            "is_layer": false,
            "is_reply_layer": false,
            "nickname": "嗯呗",
            "reply_list": [
              
            ]
          }
        ],
        "content": "因毕业出售自用人体工学椅西昊M57C\n只用了一年半，没有故障且比较新\n购买信息与现状分别如下图所示\n售价500元，目前位于科B，随时可看可提\n有意者请添加微信：jr793665605",
        "finish_status": "10",
        "ftd": "0",
        "headimgurl": "//b1.cdn.zanao.com/upload/2025/12/24/th4p4jaftrtyq6p.jpeg@!sm_w100_h100",
        "hongbao_price": "0",
        "hot_val": 0,
        "img_paths": [
          "upload/2025/12/24/g6de9gdtdsrfhpr.jpg",
          "upload/2025/12/24/kneskzp8zcsxb3w.jpg"
        ],
        "is_author": false,
        "is_forbid": "0",
        "l_count": "0",
        "l_has": false,
        "medal_exp_2023": "",
        "need_pay": false,
        "nickname": "啦啦啦625",
        "notice_ignore": false,
        "p_time": "1766583381",
        "post_time": "32分钟前",
        "report_status": "0",
        "reward_price": 2,
        "sign": "lmBpaZyaaHGWkw==",
        "thread_id": "2048967931",
        "title": "出售自用人体工学椅",
        "top_time": 0,
        "user_level": 1,
        "user_level_title": "幼儿园",
        "view_count": 196,
        "view_show": true
      }
    ],
    "tag_show": true,
    "user_cert_valid": false
  },
  "errmsg": "",
  "errno": 0
}
```

### 2.热贴

`/hot?school=scu`

获取热贴的10条帖子

返回示例：

```json
{
  "data": {
    "list": [
      {
        "c_count": "31",
        "cate_id": "108",
        "cate_name": "校园趣事",
        "color_show_info": {
          
        },
        "content": "这样想熬夜的、想复习的、想睡觉的、想吃宵夜的、想赶ddl的、想选课的就都有时间了😋",
        "finish_status": "10",
        "headimgurl": "//b1.cdn.zanao.com/upload/2025/11/20/bleeajkjdzgcuqf.jpg@!sm_w100_h100",
        "hongbao_price": "0",
        "hot_rank": 1,
        "hot_rank_show": "01",
        "hot_text": "7.1K",
        "hot_val": "165",
        "img_paths": [
          
        ],
        "l_count": "151",
        "l_has": false,
        "nickname": "吾乃天之饺子",
        "notice_ignore": false,
        "post_time": "23小时前",
        "report_status": "0",
        "thread_id": "2048595249",
        "title": "夜晚就应该增加到一百个小时",
        "user_level": 6,
        "user_level_title": "硕士生",
        "view_count": 7101,
        "view_show": true
      }
    ]
  },
  "errmsg": "",
  "errno": 0
}
```

