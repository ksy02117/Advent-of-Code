package tree

type File struct {
	Name string
	Size int
}

type Folder struct {
	Name    string
	Parent  *Folder
	Folders []Folder
	Files   []File
	Size    int
}

func (n *Folder) AddFolder(name string) {
	for _, folder := range n.Folders {
		if folder.Name == name {
			panic("directory already exists")
		}
	}

	n.Folders = append(n.Folders, Folder{
		Name:    name,
		Parent:  n,
		Folders: make([]Folder, 0),
		Files:   make([]File, 0),
		Size:    0,
	})
}

func (n *Folder) AddFile(name string, size int) {
	for _, file := range n.Files {
		if file.Name == name {
			panic("file already exists")
		}
	}

	n.Files = append(n.Files, File{
		Name: name,
		Size: size,
	})
}

func (n *Folder) MoveTo(name string) *Folder {
	for i := range n.Folders {
		if n.Folders[i].Name == name {
			return &n.Folders[i]
		}
	}
	panic("directory not found")
}

func (n *Folder) MoveToParent() *Folder {
	return n.Parent
}

func (n *Folder) CalculateSize() {
	for i := range n.Folders {
		n.Folders[i].CalculateSize()
	}

	for _, file := range n.Files {
		n.Size += file.Size
	}

	for _, folder := range n.Folders {
		n.Size += folder.Size
	}
}
