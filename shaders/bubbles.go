package shaders

import "github.com/Noofbiz/pixelshader"

var BubbleShader = &pixelshader.PixelShader{FragShader: `
#ifdef GL_ES
  #define LOWP lowp
  precision mediump float;
  #else
  #define LOWP
  #endif

uniform vec2 u_resolution;  // Canvas size (width,height)
uniform vec2 u_mouse;       // mouse position in screen pixels
uniform float u_time;       // Time in seconds since load
uniform sampler2D u_tex0;   // Drawable Tex0
uniform sampler2D u_tex1;   // Drawable Tex1
uniform sampler2D u_tex2;   // Drawable Tex2

void main()
{
	vec2 uv = -1.0 + 2.0*gl_FragCoord.xy / u_resolution.xy;
	uv.x *=  u_resolution.x / u_resolution.y;
	vec2 ms = (u_mouse.xy / u_resolution.xy) * 0.01;
    
    // background	 
	vec3 color = vec3(0.9 + 0.1*uv.y);

    // bubbles	
	for( int i=0; i<30; i++ )
	{
        // bubble seeds
		float pha =      sin(float(i)*546.13+1.0)*0.5 + 0.5;
		float siz = pow( sin(float(i)*651.74+5.0)*0.5 + 0.5, 4.0 );
		float pox =      sin(float(i)*321.55+4.1) * u_resolution.x / u_resolution.y;

        // buble size, position and color
		float rad = 0.2 + 0.5*siz;
		vec2  pos = vec2( pox, -1.0-rad + (2.0+2.0*rad)*mod(pha+0.1*u_time*(0.1+0.1*siz),1.0));
		float distToMs = length(pos - ms);
        pos *= length(pos - (ms * 1.5-0.5));
        float dis = length( uv - pos );
        
        
		vec3  col = mix( vec3(0.34,0.6,0.0), vec3(0.1,0.4,0.8), 0.5+0.5*sin(float(i)*1.2+1.9));
		    col+= 8.0*smoothstep( rad*0.95, rad, dis );
		
        // render
		float f = length(uv-pos)/rad;
		f = sqrt(clamp(1.0-f*f,0.0,1.0));
		color -= col.zyx *(1.0-smoothstep( rad*0.95, rad, dis )) * f;
	}

    // vigneting	
	color *= sqrt(1.5-0.5*length(uv));

	gl_FragColor = vec4(color,1.0);	
}
`}
