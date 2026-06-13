# Docker RUN

```
RUN set -u
RUN echo $a
# result: won't fail
```

```
RUN set -u; echo $a
# result: fail
```
