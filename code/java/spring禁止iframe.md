# spring 禁止 iframe

```java
@EnableWebSecurity
@Configuration
public class WebSecurityConfig extends DefaultWebSecurityConfigurer {

    @Override
    protected void configure(HttpSecurity http) throws Exception {
        super.configure(http);
        // X-Frame-Options 禁止
        http.headers().frameOptions().disable();
    }
}
```
