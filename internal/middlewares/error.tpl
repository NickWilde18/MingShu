<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.HTTPStatus}} - {{.TitleZh}} / {{.TitleEn}}</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", "Helvetica Neue", Arial, "Noto Sans", sans-serif;
            background: #f5f7fa;
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 20px;
            color: #333;
        }
        
        .error-container {
            background: white;
            border-radius: 8px;
            box-shadow: 0 2px 12px rgba(0,0,0,0.08);
            max-width: 800px;
            width: 100%;
            overflow: hidden;
            animation: slideIn 0.5s ease-out;
            border: 1px solid #e1e8ed;
        }
        
        @keyframes slideIn {
            from {
                opacity: 0;
                transform: translateY(-30px);
            }
            to {
                opacity: 1;
                transform: translateY(0);
            }
        }
        
        .error-header {
            background: #1a1a2e;
            color: white;
            padding: 25px 30px;
            text-align: center;
            border-bottom: 3px solid #0f3460;
        }
        
        .error-code {
            font-size: 56px;
            font-weight: 700;
            margin-bottom: 8px;
            text-shadow: 2px 2px 4px rgba(0,0,0,0.2);
        }
        
        .error-title {
            font-size: 20px;
            font-weight: 500;
            opacity: 0.95;
            line-height: 1.5;
        }
        
        .error-title .zh {
            display: block;
        }
        
        .error-title .en {
            display: block;
            font-size: 16px;
            opacity: 0.85;
            margin-top: 4px;
        }
        
        .error-body {
            padding: 25px 30px;
        }
        
        .error-message {
            font-size: 15px;
            line-height: 1.6;
            color: #555;
            margin-bottom: 20px;
        }
        
        .bilingual .zh {
            color: #333;
            margin-bottom: 8px;
        }
        
        .bilingual .en {
            color: #666;
            font-size: 14px;
        }
        
        .custom-message {
            background: #fff3cd;
            border-left: 4px solid #ffc107;
            padding: 12px 16px;
            margin: 15px 0;
            border-radius: 4px;
            color: #856404;
            font-size: 14px;
        }
        
        .error-reason {
            background: #f8f9fa;
            border-left: 4px solid #dc3545;
            padding: 12px 16px;
            margin: 15px 0;
            border-radius: 4px;
        }
        
        .reason-label {
            font-weight: 600;
            color: #dc3545;
            margin-bottom: 8px;
            font-size: 13px;
        }
        
        .reason-text {
            color: #666;
            font-size: 13px;
            line-height: 1.5;
            font-family: "Monaco", "Menlo", "Consolas", monospace;
            word-break: break-word;
        }
        
        .suggestion-box {
            background: #f8f9fa;
            border-left: 4px solid #0f3460;
            padding: 15px;
            margin: 20px 0;
            border-radius: 4px;
        }
        
        .suggestion-title {
            font-weight: 600;
            color: #0f3460;
            margin-bottom: 10px;
            font-size: 14px;
        }
        
        .suggestion-content {
            line-height: 1.6;
        }
        
        .info-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
            gap: 15px;
            margin: 20px 0;
            padding: 15px;
            background: #f8f9fa;
            border-radius: 6px;
        }
        
        .info-item {
            display: flex;
            flex-direction: column;
        }
        
        .info-label {
            font-size: 11px;
            color: #999;
            margin-bottom: 4px;
            text-transform: uppercase;
            letter-spacing: 0.5px;
        }
        
        .info-value {
            font-size: 13px;
            color: #333;
            font-weight: 500;
            font-family: "Monaco", "Menlo", "Consolas", monospace;
            word-break: break-all;
        }
        
        .info-value.traceid {
            position: relative;
            display: inline-flex;
            align-items: center;
            gap: 6px;
        }
        
        .info-value.traceid > span:first-child {
            cursor: help;
            position: relative;
        }
        
        .info-value.traceid > span:first-child:hover::after {
            content: attr(title);
            position: absolute;
            bottom: 100%;
            left: 0;
            background: rgba(255, 255, 255, 0.98);
            color: #333;
            padding: 8px 12px;
            border-radius: 4px;
            font-size: 12px;
            white-space: nowrap;
            z-index: 1000;
            margin-bottom: 5px;
            box-shadow: 0 2px 12px rgba(0,0,0,0.15);
            border: 1px solid #e0e0e0;
        }
        
        .info-value.traceid > span:first-child:hover::before {
            content: '';
            position: absolute;
            bottom: 100%;
            left: 10px;
            border: 5px solid transparent;
            border-top-color: rgba(255, 255, 255, 0.98);
            margin-bottom: -5px;
            z-index: 1001;
        }
        
        .copy-icon {
            display: inline-flex;
            align-items: center;
            justify-content: center;
            width: 20px;
            height: 20px;
            cursor: pointer;
            border-radius: 3px;
            transition: all 0.2s;
            position: relative;
        }
        
        .copy-icon:hover {
            background: #e9ecef;
        }
        
        .copy-icon::after {
            content: '';
            display: inline-block;
            width: 14px;
            height: 14px;
            background-image: url('data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="%230f3460" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></svg>');
            background-size: contain;
            background-repeat: no-repeat;
            background-position: center;
            opacity: 0.5;
            transition: opacity 0.2s;
        }
        
        .copy-icon:hover::after {
            opacity: 1;
        }
        
        .copy-tooltip {
            position: absolute;
            bottom: 100%;
            left: 50%;
            transform: translateX(-50%);
            background: rgba(15, 52, 96, 0.9);
            color: white;
            padding: 6px 12px;
            border-radius: 4px;
            font-size: 12px;
            white-space: nowrap;
            margin-bottom: 5px;
            opacity: 0;
            pointer-events: none;
            transition: opacity 0.2s;
            box-shadow: 0 2px 8px rgba(0,0,0,0.1);
        }
        
        .copy-tooltip.show {
            opacity: 1;
        }
        
        .copy-tooltip::after {
            content: '';
            position: absolute;
            top: 100%;
            left: 50%;
            transform: translateX(-50%);
            border: 5px solid transparent;
            border-top-color: rgba(15, 52, 96, 0.9);
        }
        
        .detail-section {
            margin: 20px 0;
        }
        
        .detail-toggle {
            width: 100%;
            padding: 12px;
            background: #f8f9fa;
            border: 1px solid #e0e0e0;
            border-radius: 4px;
            cursor: pointer;
            font-size: 13px;
            font-weight: 500;
            color: #0f3460;
            display: flex;
            justify-content: space-between;
            align-items: center;
            transition: all 0.3s;
        }
        
        .detail-toggle:hover {
            background: #e9ecef;
        }
        
        .detail-content {
            margin-top: 8px;
            padding: 15px;
            background: #f8f9fa;
            border-radius: 4px;
            border: 1px solid #e0e0e0;
            max-height: 200px;
            overflow: auto;
        }
        
        .detail-content pre {
            margin: 0;
            font-size: 11px;
            line-height: 1.4;
            color: #666;
            white-space: pre-wrap;
            word-wrap: break-word;
            font-family: "Monaco", "Menlo", "Consolas", monospace;
        }
        
        .action-buttons {
            display: flex;
            gap: 12px;
            justify-content: center;
            flex-wrap: wrap;
            margin-top: 25px;
        }
        
        .btn {
            padding: 10px 24px;
            border-radius: 4px;
            text-decoration: none;
            font-weight: 500;
            font-size: 14px;
            transition: all 0.3s;
            border: none;
            cursor: pointer;
            display: inline-flex;
            align-items: center;
            gap: 6px;
        }
        
        .btn-primary {
            background: #0f3460;
            color: white;
            box-shadow: 0 2px 8px rgba(15, 52, 96, 0.2);
        }
        
        .btn-primary:hover {
            background: #16213e;
            box-shadow: 0 4px 12px rgba(15, 52, 96, 0.3);
        }
        
        .btn-secondary {
            background: white;
            color: #0f3460;
            border: 2px solid #0f3460;
        }
        
        .btn-secondary:hover {
            background: #0f3460;
            color: white;
        }
        
        .error-footer {
            background: #f8f9fa;
            padding: 15px 30px;
            text-align: center;
            font-size: 13px;
            color: #999;
            border-top: 1px solid #e0e0e0;
        }
        
        .error-footer .zh {
            display: block;
        }
        
        .error-footer .en {
            display: block;
            font-size: 11px;
            margin-top: 4px;
        }
        
        @media (max-width: 768px) {
            .error-code {
                font-size: 48px;
            }
            
            .error-title {
                font-size: 20px;
            }
            
            .error-title .en {
                font-size: 16px;
            }
            
            .error-body {
                padding: 30px 20px;
            }
            
            .info-grid {
                grid-template-columns: 1fr;
                gap: 15px;
            }
            
            .action-buttons {
                flex-direction: column;
            }
            
            .btn {
                width: 100%;
                justify-content: center;
            }
        }
    </style>
</head>
<body>
    <div class="error-container">
        <div class="error-header">
            <div class="error-code">{{.HTTPStatus}}</div>
            <div class="error-title">
                <span class="zh">{{.TitleZh}}</span>
                <span class="en">{{.TitleEn}}</span>
            </div>
        </div>
        
        <div class="error-body">
            <div class="error-message bilingual">
                <div class="zh">{{.MessageZh}}</div>
                <div class="en">{{.MessageEn}}</div>
            </div>
            
            {{if .CustomMsg}}
            <div class="custom-message">
                {{.CustomMsg}}
            </div>
            {{end}}
            
            {{if .SuggestionZh}}
            <div class="suggestion-box">
                <div class="suggestion-title">建议 / Suggestions</div>
                <div class="suggestion-content bilingual">
                    <div class="zh">{{.SuggestionZh}}</div>
                    <div class="en">{{.SuggestionEn}}</div>
                </div>
            </div>
            {{end}}
            
            {{if and .Detail (not .ShowDetail)}}
            <div class="error-reason">
                <div class="reason-label">错误原因 / Error Reason</div>
                <div class="reason-text">{{.ReasonText}}</div>
            </div>
            {{end}}
            
            <div class="info-grid">
                <div class="info-item">
                    <div class="info-label">追踪ID / Trace ID</div>
                    <div class="info-value traceid">
                        <span title="{{.TraceID}}">{{.TraceIDShort}}</span>
                        <span class="copy-icon" data-full-id="{{.TraceID}}" onclick="copyTraceID(this)">
                            <span class="copy-tooltip" id="copy-tooltip">已复制!</span>
                        </span>
                    </div>
                </div>
                <div class="info-item">
                    <div class="info-label">错误代码 / Error Code</div>
                    <div class="info-value">#{{.ErrorCode}}</div>
                </div>
                <div class="info-item">
                    <div class="info-label">错误时间 / Time</div>
                    <div class="info-value">{{.Timestamp}}</div>
                </div>
            </div>
            
            {{if and .ShowDetail .Detail}}
            <div class="detail-section">
                <button class="detail-toggle" onclick="toggleDetail()">
                    <span id="toggle-text">显示详情 Show Details</span>
                    <span id="toggle-icon">▼</span>
                </button>
                <div class="detail-content" id="detail-content" style="display: none;">
                    <pre>{{.Detail}}</pre>
                </div>
            </div>
            {{end}}
            
            <div class="action-buttons">
                {{if and .ButtonLeft .ButtonLeftJS}}
                <a onclick="{{.ButtonLeftJS}}" class="btn btn-primary">
                    {{.ButtonLeft}}
                </a>
                {{end}}
                {{if and .ButtonRight .ButtonRightJS}}
                <button onclick="{{.ButtonRightJS}}" class="btn btn-secondary">
                    {{.ButtonRight}}
                </button>
                {{end}}
            </div>
        </div>
        
        <div class="error-footer">
            <span class="zh">香港中文大学（深圳）GPT服务平台</span>
            <span class="en">CUHK-Shenzhen GPT Services Platform</span>
        </div>
    </div>
    
    <script>
        function toggleDetail() {
            const content = document.getElementById('detail-content');
            const text = document.getElementById('toggle-text');
            const icon = document.getElementById('toggle-icon');
            
            if (content.style.display === 'none') {
                content.style.display = 'block';
                text.textContent = '隐藏详情 Hide Details';
                icon.textContent = '▲';
            } else {
                content.style.display = 'none';
                text.textContent = '显示详情 Show Details';
                icon.textContent = '▼';
            }
        }
        
        function copyTraceID(element) {
            const fullID = element.getAttribute('data-full-id');
            
            // 复制到剪贴板
            navigator.clipboard.writeText(fullID).then(function() {
                // 显示提示
                const tooltip = document.getElementById('copy-tooltip');
                tooltip.classList.add('show');
                
                // 2秒后隐藏
                setTimeout(function() {
                    tooltip.classList.remove('show');
                }, 2000);
            }).catch(function(err) {
                // 兼容旧浏览器的方法
                const textArea = document.createElement('textarea');
                textArea.value = fullID;
                textArea.style.position = 'fixed';
                textArea.style.left = '-999999px';
                document.body.appendChild(textArea);
                textArea.select();
                
                try {
                    document.execCommand('copy');
                    const tooltip = document.getElementById('copy-tooltip');
                    tooltip.classList.add('show');
                    setTimeout(function() {
                        tooltip.classList.remove('show');
                    }, 2000);
                } catch (err) {
                    console.error('复制失败:', err);
                }
                
                document.body.removeChild(textArea);
            });
        }
    </script>
    {{if .CustomJS}}
    <script>
        {{.CustomJS}}
    </script>
    {{end}}
</body>
</html>

