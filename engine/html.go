package engine

const baseTemplate = `<!DOCTYPE html>
<html lang="zh-cn">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>POC-Runner</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #121212;
            color: #ffffff;
            margin: 0;
            padding: 20px;
        }

        .container {
            max-width: 1200px;
            margin: auto;
            padding: 20px;
        }

        h1 {
            text-align: center;
            font-size: 2.5em;
            margin-bottom: 30px;
            color: #00bcd4;
            text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.7);
        }

        .report-item {
            background-color: #1e1e1e;
            margin-bottom: 20px;
            border-radius: 8px;
            padding: 15px;
            transition: background-color 0.3s;
        }

        .report-item:hover {
            background-color: #292929;
        }

        .report-header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            cursor: pointer;
        }

        .report-header h2 {
            margin: 0;
        }

        .report-header span {
            color: #888;
        }

        .rules {
            display: none;
            margin-top: 15px;
        }

        .rules .rule-item {
            background-color: #2a2a2a;
            padding: 10px;
            border-radius: 5px;
            margin-bottom: 10px;
        }

        .rule-item pre {
            background-color: #333;
            padding: 10px;
            border-radius: 5px;
            overflow: auto;
            white-space: pre-wrap;
            word-wrap: break-word; 
        }
    </style>
</head>
<body>
%s
<script>
    function toggleRules(headerElement) {
        const rules = headerElement.nextElementSibling;
        if (rules.style.display === 'block') {
            rules.style.display = 'none';
        } else {
            rules.style.display = 'block';
        }
    }
</script>
</body>
</html>`

const dataTemplate = `<div class="container">
    <h1>POC-Runner Vulnerability Report</h1>
    <div class="report-item">
        <div class="report-header" onclick="toggleRules(this)">
            <h2>%s</h2>
            <span>Target: %s</span>
            <span>Time: %s</span>
        </div>
        <div class="rules">
            %s
        </div>
    </div>
</div>`

const ruleTemplate = `<div class="rule-item">
    <h3>%s</h3>
    <p><strong>Request </strong></p>
    <pre>%s</pre>
    <p><strong>Response </strong></p>
    <pre>%s</pre>
    <p><strong>Result </strong> %t</p>
</div>`
