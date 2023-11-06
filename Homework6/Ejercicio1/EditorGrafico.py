class Document:
    def __init__(self, title):
        self.title = title
        self.pages = []

    def add_page(self, page):
        self.pages.append(page)

    def remove_page(self, page):
        if page in self.pages:
            self.pages.remove(page)

    def display(self):
        print(f"Document Title: {self.title}")
        for page in self.pages:
            page.display()

class Page:
    def __init__(self):
        self.elements = []

    def add_element(self, element):
        self.elements.append(element)

    def remove_element(self, element):
        if element in self.elements:
            self.elements.remove(element)

    def display(self):
        print("Page Content:")
        for element in self.elements:
            element.display()

class Element:
    def __init__(self):
        pass

    def display(self):
        pass

class Text(Element):
    def __init__(self, text):
        super().__init__()
        self.text = text

    def display(self):
        print(f"Text: {self.text}")

class Shape(Element):
    def __init__(self, shape_type):
        super().__init__()
        self.shape_type = shape_type

    def display(self):
        print(f"Shape: {self.shape_type}")

class Group(Element):
    def __init__(self):
        super().__init__()
        self.elements = []

    def add_element(self, element):
        self.elements.append(element)

    def remove_element(self, element):
        if element in self.elements:
            self.elements.remove(element)

    def display(self):
        print("Group:")
        for element in self.elements:
            element.display()


# Ejemplo de uso
doc = Document("My Document")

page1 = Page()
page1.add_element(Text("Hello, World!"))
page1.add_element(Shape("Circle"))

page2 = Page()
group = Group()
group.add_element(Text("Grouped Text"))
group.add_element(Shape("Rectangle"))
page2.add_element(group)
page2.add_element(Shape("Line"))

doc.add_page(page1)
doc.add_page(page2)

doc.display()
