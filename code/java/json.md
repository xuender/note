# JSON 转换

## jackson

依赖

```xml
<dependency>
    <groupId>com.fasterxml.jackson.core</groupId>
    <artifactId>jackson-core</artifactId>
    <version>2.11.0</version>
</dependency>
<dependency>
    <groupId>com.fasterxml.jackson.core</groupId>
    <artifactId>jackson-databind</artifactId>
    <version>2.11.0</version>
</dependency>
```

转换

```java
public final class JsonUtils {
  private static final ObjectMapper mapper = new ObjectMapper();
  private static final JsonFactory jsonFactory = new JsonFactory();

  public static String toJSONString(Object obj) throws IOException {
    StringWriter sw = new StringWriter();
    JsonGenerator gen = jsonFactory.createGenerator(sw);
    mapper.writeValue(gen, obj);
    gen.close();
    return sw.toString();
  }

  public static <T> T parseObject(String jsonStr, Class<T> objClass)
    throws IOException {
    return mapper.readValue(jsonStr, objClass);
  }

  public static <T> T parseObject(String jsonStr, TypeReference<T> valueTypeRef)
    throws IOException {
    return mapper.readValue(jsonStr, valueTypeRef);
  }
}
```

## FastJson

bug太多，代码质量不好，别用
