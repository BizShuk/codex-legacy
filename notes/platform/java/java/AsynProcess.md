# Async


### Annotation
@EnableAsync
@Async


### @EnableAsync
1. Put @EnableAsync in configuration file
2. Implement `AsyncConfigurer`
3. Override `getAsyncExecutor` and `getAsyncUncaughtExceptionHandler`

ps: If there is noe Executor specified, `SimpleAsyncTaskExecutor` will be used. 

```
@Override
  public Executor getAsyncExecutor() {
    ThreadPoolTaskExecutor executor = new ThreadPoolTaskExecutor();
    executor.setCorePoolSize(coreSize);
    executor.setMaxPoolSize(maxSize);
    executor.setQueueCapacity(queueCapacity);
    executor.setThreadNamePrefix("worker-exec-");
    executor.setRejectedExecutionHandler(new ThreadPoolExecutor.CallerRunsPolicy());
    return executor;
  }
  @Override
  public AsyncUncaughtExceptionHandler getAsyncUncaughtExceptionHandler() {
    return (ex, method, params) -> {
      Class<?> targetClass = method.getDeclaringClass();
      Logger logger = LoggerFactory.getLogger(targetClass);
      logger.error(ex.getMessage(), ex);
    };
  }

```

### RejectedExecutionHandler vs AsyncUncaughtExceptionHandler
RejectedExecutionHandler is for that the task executor doesn't accpet the task to be executed. If the task executor is not shut down, it should be able to put back to task executor queue again. Otherwise, it should be logged.

AsyncUncaughtExceptionHandler is for that the async method throw out an exception which is not handled. This one can recall, put it in queue, or log it.

### @Async
value for qualifier for specified async resource (task executor)

Either void or Future return type.




### Additional
- AysncConfigurationSelector
- ProxyAsyncConfiguration
- AsynExecutionInterceptor extends AsyncExecutionAspectSupport
