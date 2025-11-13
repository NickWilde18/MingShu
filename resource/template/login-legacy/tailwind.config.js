/** @type {import('tailwindcss').Config} */
export default {
  content: ["./src/**/*.{html,js}", "./js/**/*.js"],
  darkMode: 'class',
  theme: {
    screens: {
      'sm': '640px',
      'md': '768px', 
      'lg': '1024px',
      'xl': '1280px',
      '2xl': '1536px',
    },
    extend: {
      colors: {
        // 主品牌色 - 深紫色系
        primary: {
          50: '#f5f3ff',
          100: '#ede9fe', 
          200: '#ddd6fe',
          300: '#c4b5fd',
          400: '#a78bfa',
          500: '#8b5cf6',
          600: '#7c3aed',
          700: '#6d28d9',
          800: '#5b21b6',
          900: '#4c1d95',
          950: '#2e1065'
        },
        // 渐变色系 - 紫色到蓝色
        gradient: {
          start: '#8b5cf6', // 紫色
          middle: '#3b82f6', // 蓝色
          end: '#06b6d4'     // 青色
        },
        // 背景色系
        background: {
          light: '#fafafa',
          dark: '#0f0f23',
          card: {
            light: 'rgba(255, 255, 255, 0.85)',
            dark: 'rgba(30, 30, 60, 0.85)'
          }
        },
        // 文本色系
        text: {
          primary: {
            light: 'rgba(190, 151, 239, 1)',
            dark: '#f9fafb'
          },
          secondary: {
            light: '#6b7280',
            dark: '#d1d5db'
          },
          muted: {
            light: '#9ca3af',
            dark: '#9ca3af'
          },
          switch: {
            light: 'rgba(202, 202, 218, 1)',
            dark: 'rgba(89, 70, 207, 1)'
          }
        },
        // 边框色系
        border: {
          light: 'rgba(255, 255, 255, 0.6)',
          dark: 'rgba(75, 85, 180, 0.6)'
        }
      },
      backgroundImage: {
        // 主按钮渐变
        'gradient-primary': 'linear-gradient(135deg, #8b5cf6 0%, #3b82f6 50%, #06b6d4 100%)',
        // 卡片光晕效果
        'card-glow': 'conic-gradient(from 90deg, rgba(139, 92, 246, 0.3), rgba(59, 130, 246, 0.2), rgba(6, 182, 212, 0.2), rgba(139, 92, 246, 0.3))',
        // 背景渐变 - 浅色模式
        'bg-light': 'linear-gradient(135deg, #f0e9ff 0%, #e7eafe 60%, #d1f8fe 100%)',
        // 背景渐变 - 深色模式  
        'bg-dark': 'linear-gradient(135deg, #0f0f23 0%, #1e1e3f 65%, #142155 100%)'
      },
      boxShadow: {
        'card': '0 25px 50px -12px rgba(0, 0, 0, 0.25)',
        'card-dark': '0 25px 50px -12px rgba(0, 0, 0, 0.5)',
        'glow': '0 0 20px rgba(139, 92, 246, 0.3)',
        'glow-dark': '0 0 20px rgba(139, 92, 246, 0.2)'
      },
      backdropBlur: {
        'xs': '2px',
      }
    },
  },
  plugins: [],
} 