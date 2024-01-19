# 检查word文档中是否有重复的指定字符串
# TODO：我python烂，先这样吧，之后再说
import re
from docx import Document


def find_duplicate_fields(doc_path):
    document = Document(doc_path)

    # 定义要查找的字段
    fields_to_check = ['name', 'email', 'phone']

    # 创建字典来存储字段和它们的值
    fields_values = {field: set() for field in fields_to_check}

    for paragraph in document.paragraphs:
        for field in fields_to_check:
            # 使用正则表达式查找字段和值
            pattern = re.compile(rf'{field}:\s*(.*)', re.IGNORECASE)
            match = pattern.search(paragraph.text)

            if match:
                value = match.group(1).strip()
                # 如果值已经存在，表示文档中存在作者信息
                if value in fields_values[field]:
                    with open(doc_path+'.invalid', 'w'):
                        pass
                    return
                else:
                    fields_values[field].add(value)


if __name__ == "__main__":
    # 指定Word文档的路径
    word_document_path = 'path/to/your/document.docx'

    find_duplicate_fields(word_document_path)
