# Auth server for micro service

Intro
-----

로그인 서버의 비대칭쌍키(RSA, ECDSA)를 생성하고, 시스템의 타 서버에 공개키를 할당한다.<br/> 
로그인 서버는 사용자의 올바른 로그인 요청에 서명하여 토큰을 발급한다.<br/> 
토큰은 expired time과 사용자 계정의 정보를 담고 있으며, 단방향으로 토큰으로 부터 사용자 계정의 정보를 추출하는 것은 불가능하고, <br/>
토큰이 로그인 서버에 의해 발급된 사실과 토큰의 정보가 위조되지 않았는가에 대한 검증만 가능하다.


주의 > 서명은 옵션의 조건(데드라인 등) 외에 파기할 수 있는 방법이 없습니다. 클라이언트에서 영구적이지 않은 보관을 통해주세요

option
gen = false (default : true)  신규 비밀키를 발급하지 않습니다.

