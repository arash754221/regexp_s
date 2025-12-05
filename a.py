import re

# الگوهای بهبود یافته
token_specification = [
    ("NUMBER", r"\d+(\.\d*)?"),  # عدد
    ("ASSIGN", r"="),  # علامت مساوی
    ("END", r";"),  # سمی‌کالن
    ("ID", r"[A-Za-z_]\w*"),  # شناسه
    ("OP", r"[+\-*/]"),  # عملگرها
    ("LPAREN", r"\("),  # پرانتز باز
    ("RPAREN", r"\)"),  # پرانتز بسته
    ("COLON", r":"),  # کولن
    ("COMMA", r","),  # کاما
    ("NEWLINE", r"\n"),  # پایان خط
    ("SKIP", r"[ \t]+"),  # فاصله‌ها و تب
    ("MISMATCH", r"."),  # هر چیز غیرمنتظره
]

tok_regex = "|".join("(?P<%s>%s)" % pair for pair in token_specification)

code = input("کد رو بنویس: ")

for mo in re.finditer(tok_regex, code):
    kind = mo.lastgroup
    value = mo.group()
    if kind == "NUMBER":
        value = float(value) if "." in value else int(value)
    elif kind == "ID" and value in {"if", "else", "for", "while", "range", "print"}:
        kind = "KEYWORD"
    elif kind == "SKIP" or kind == "NEWLINE":
        continue
    elif kind == "MISMATCH":
        print(f'❌ خطای لغوی در "{value}"')
        break
    print(f"{kind:10} → {value}")
