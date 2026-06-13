# Cross-Region Design

As US branch started to involve, the service needs to be extended globally. But there are some US policies and political issues between Taiwan and China [#1].

US policies, like data is belong to US if data stored in US territory, is dangering TSMC advantages in semiconductor industry. Therefore, all machine learning model should be stored in memory only. With this precondition, edging computing comes to my mind. DevOps CI/CD is also playing a big role here. But some tricky points are that no more operation people in any site. It comes from CTO, from FB, decision [#2]. The architecture becomes so skew to central code base/monitoring/model training, and edge interface (like Central Web Service, not global design).

No doubt, queue/cache are leveraged as scaling means. How to load model into specific region, issues with region delay(both side alerting)

[#1]: TSMC is Silicon Shield for Taiwan to prevent invension from China. Most of countries rely on chips built from TSMC.

[#2]: It's arbitrary decision and blindness of followers.
