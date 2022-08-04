# ncbi asn 对应的数据结构



数据结构根据ncbi公开的数据定义文件生成并进行手工处理

源数据结构文件：https://ncbi.nlm.nih.gov/data_specs/asn/NCBI_all.asn

asn文件中的ENUMERATED类型转换成golang的int64类型
asn文件中的NULL类型转换成golang的interface{}类型

数据结构说明详细查看asn定义文件中的对应结构说明

注意：
    这里的数据文件与ncbi datatool工具生成的xml数据结构并不一致。
    datatool生成的xml xsd定义文件中把结构定义本身单独生成了一个数据类型。生成的xml文件会多出很多没用的结构节点。