# ImageOptimizer

@powered by [imaging](github.com/disintegration/imaging)

### API

> GET /{image-name}

{image-name} 의 원본을 요청

> GET /{image-name}?width=123&height=123

원하는 크기로 조절된 {image-name} 을 불러

width 또는 height 중 하나만 제공할경우, 기존 이미지 사이즈의 비율에 맞춰 줄어듦.
만약 둘 다 제공할 경우, 비율을 고려하지 않은채 절대사이즈에 맞게 줄어듦.

> POST /

multipart/form-data, key: image

이미지를 업로드함. 이미지의 저장 이름은 파일이름을 가져옴.  
추후 웹 인터페이스를 제공할 예정.


### 사용 설명

1. 파일 저장 위치

   원본 : ./images/ori  
   Optimize 이미지 캐싱 : ./images/opt

   현재 캐시된 이미지를 삭제하는 정책은 아직 존재하지 않음. (빠른 시일내로 추가할 예정)

2. Resizing

   이미지의 사이즈를 조절하게 되면 모든 이미지는 png 의 형태로 제공됨. (이것도 추후 옵션을 통해 jpeg 를 지원할 예정)
   
   사이즈가 조절된 이미지는 캐싱되어 따로 저장됨. 
